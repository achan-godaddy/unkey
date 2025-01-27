import { TRPCError } from "@trpc/server";
import { z } from "zod";

import { insertAuditLogs } from "@/lib/audit";
import { and, db, eq, isNull, schema, sql } from "@/lib/db";
import { ingestAuditLogsTinybird } from "@/lib/tinybird";
import { rateLimitedProcedure, ratelimit } from "@/lib/trpc/ratelimitProcedure";
import { newId } from "@unkey/id";

export const createOverride = rateLimitedProcedure(ratelimit.create)
  .input(
    z.object({
      namespaceId: z.string(),
      identifier: z.string(),
      limit: z.number(),
      duration: z.number(),
      async: z.boolean().optional(),
    }),
  )
  .mutation(async ({ input, ctx }) => {
    const namespace = await db.query.ratelimitNamespaces
      .findFirst({
        where: (table, { and, eq, isNull }) =>
          and(eq(table.id, input.namespaceId), isNull(table.deletedAt)),
        with: {
          workspace: {
            columns: {
              id: true,
              tenantId: true,
              features: true,
            },
          },
        },
      })
      .catch((_err) => {
        throw new TRPCError({
          code: "INTERNAL_SERVER_ERROR",
          message:
            "We are unable to create an override for this namespace. Please contact support using support@unkey.dev",
        });
      });
    if (!namespace || namespace.workspace.tenantId !== ctx.tenant.id) {
      throw new TRPCError({
        code: "NOT_FOUND",
        message:
          "We are unable to find the correct namespace. Please contact support using support@unkey.dev.",
      });
    }

    const id = newId("ratelimitOverride");
    await db
      .transaction(async (tx) => {
        const existing = await tx
          .select({ count: sql`count(*)` })
          .from(schema.ratelimitOverrides)
          .where(
            and(
              eq(schema.ratelimitOverrides.namespaceId, namespace.id),
              isNull(schema.ratelimitOverrides.deletedAt),
            ),
          )
          .then((res) => Number(res.at(0)?.count ?? 0));
        const max =
          typeof namespace.workspace.features.ratelimitOverrides === "number"
            ? namespace.workspace.features.ratelimitOverrides
            : 5;
        if (existing >= max) {
          throw new TRPCError({
            code: "FORBIDDEN",
            message: `A plan Upgrade is required, you can only override ${max} identifiers.`,
          });
        }

        await tx.insert(schema.ratelimitOverrides).values({
          workspaceId: namespace.workspace.id,
          namespaceId: namespace.id,
          identifier: input.identifier,
          id,
          limit: input.limit,
          duration: input.duration,
          createdAt: new Date(),
          async: input.async,
        });
        await insertAuditLogs(tx, {
          workspaceId: namespace.workspace.id,
          actor: {
            type: "user",
            id: ctx.user.id,
          },
          event: "ratelimitOverride.create",
          description: `Created ${input.identifier}`,
          resources: [
            {
              type: "ratelimitNamespace",
              id: input.namespaceId,
            },
            {
              type: "ratelimitOverride",
              id,
            },
          ],
          context: {
            location: ctx.audit.location,
            userAgent: ctx.audit.userAgent,
          },
        });
      })
      .catch((_err) => {
        throw new TRPCError({
          code: "INTERNAL_SERVER_ERROR",
          message:
            "We are unable to create the override. Please contact support using support@unkey.dev",
        });
      });

    await ingestAuditLogsTinybird({
      workspaceId: namespace.workspace.id,
      actor: {
        type: "user",
        id: ctx.user.id,
      },
      event: "ratelimitOverride.create",
      description: `Created ${input.identifier}`,
      resources: [
        {
          type: "ratelimitNamespace",
          id: input.namespaceId,
        },
        {
          type: "ratelimitOverride",
          id,
        },
      ],
      context: {
        location: ctx.audit.location,
        userAgent: ctx.audit.userAgent,
      },
    });

    return {
      id,
    };
  });
