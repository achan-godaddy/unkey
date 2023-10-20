// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/keys/v1/keys.proto

package keysv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RatelimitType int32

const (
	RatelimitType_RATELIMIT_TYPE_UNSPECIFIED RatelimitType = 0
	RatelimitType_RATELIMIT_TYPE_FAST        RatelimitType = 1
	RatelimitType_RATELIMIT_TYPE_CONSISTENT  RatelimitType = 2
)

// Enum value maps for RatelimitType.
var (
	RatelimitType_name = map[int32]string{
		0: "RATELIMIT_TYPE_UNSPECIFIED",
		1: "RATELIMIT_TYPE_FAST",
		2: "RATELIMIT_TYPE_CONSISTENT",
	}
	RatelimitType_value = map[string]int32{
		"RATELIMIT_TYPE_UNSPECIFIED": 0,
		"RATELIMIT_TYPE_FAST":        1,
		"RATELIMIT_TYPE_CONSISTENT":  2,
	}
)

func (x RatelimitType) Enum() *RatelimitType {
	p := new(RatelimitType)
	*p = x
	return p
}

func (x RatelimitType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RatelimitType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_keys_v1_keys_proto_enumTypes[0].Descriptor()
}

func (RatelimitType) Type() protoreflect.EnumType {
	return &file_proto_keys_v1_keys_proto_enumTypes[0]
}

func (x RatelimitType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RatelimitType.Descriptor instead.
func (RatelimitType) EnumDescriptor() ([]byte, []int) {
	return file_proto_keys_v1_keys_proto_rawDescGZIP(), []int{0}
}

type Ratelimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           RatelimitType `protobuf:"varint,1,opt,name=type,proto3,enum=proto.keys.v1.RatelimitType" json:"type,omitempty"`
	Limit          int32         `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	RefillRate     int32         `protobuf:"varint,3,opt,name=refill_rate,json=refillRate,proto3" json:"refill_rate,omitempty"`
	RefillInterval int32         `protobuf:"varint,4,opt,name=refill_interval,json=refillInterval,proto3" json:"refill_interval,omitempty"`
}

func (x *Ratelimit) Reset() {
	*x = Ratelimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_keys_v1_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ratelimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ratelimit) ProtoMessage() {}

func (x *Ratelimit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_keys_v1_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ratelimit.ProtoReflect.Descriptor instead.
func (*Ratelimit) Descriptor() ([]byte, []int) {
	return file_proto_keys_v1_keys_proto_rawDescGZIP(), []int{0}
}

func (x *Ratelimit) GetType() RatelimitType {
	if x != nil {
		return x.Type
	}
	return RatelimitType_RATELIMIT_TYPE_UNSPECIFIED
}

func (x *Ratelimit) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Ratelimit) GetRefillRate() int32 {
	if x != nil {
		return x.RefillRate
	}
	return 0
}

func (x *Ratelimit) GetRefillInterval() int32 {
	if x != nil {
		return x.RefillInterval
	}
	return 0
}

type CreateKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId string  `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	KeyAuthId   string  `protobuf:"bytes,2,opt,name=key_auth_id,json=keyAuthId,proto3" json:"key_auth_id,omitempty"`
	Prefix      *string `protobuf:"bytes,3,opt,name=prefix,proto3,oneof" json:"prefix,omitempty"`
	Name        *string `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	ByteLength  *int32  `protobuf:"varint,5,opt,name=byte_length,json=byteLength,proto3,oneof" json:"byte_length,omitempty"`
	OwnerId     *string `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3,oneof" json:"owner_id,omitempty"`
	Meta        *string `protobuf:"bytes,7,opt,name=meta,proto3,oneof" json:"meta,omitempty"`
	// Unix milliseconds
	Expires   *int64     `protobuf:"varint,8,opt,name=expires,proto3,oneof" json:"expires,omitempty"`
	Ratelimit *Ratelimit `protobuf:"bytes,9,opt,name=ratelimit,proto3,oneof" json:"ratelimit,omitempty"`
	// for_workspace_id marks this key as a root key.
	// workspace_id is set to our internal Unkey workspace while
	// for_workspace_id is set to the workspace id of the user.
	//
	// when verifying a root key, we return this for_workspace_id to authenticate the user.
	ForWorkspaceId *string `protobuf:"bytes,10,opt,name=for_workspace_id,json=forWorkspaceId,proto3,oneof" json:"for_workspace_id,omitempty"`
	// How often this key may be used
	// `undefined`, `0` or negative to disable
	Remaining *int32 `protobuf:"varint,11,opt,name=remaining,proto3,oneof" json:"remaining,omitempty"`
}

func (x *CreateKeyRequest) Reset() {
	*x = CreateKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_keys_v1_keys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeyRequest) ProtoMessage() {}

func (x *CreateKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_keys_v1_keys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_keys_v1_keys_proto_rawDescGZIP(), []int{1}
}

func (x *CreateKeyRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CreateKeyRequest) GetKeyAuthId() string {
	if x != nil {
		return x.KeyAuthId
	}
	return ""
}

func (x *CreateKeyRequest) GetPrefix() string {
	if x != nil && x.Prefix != nil {
		return *x.Prefix
	}
	return ""
}

func (x *CreateKeyRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CreateKeyRequest) GetByteLength() int32 {
	if x != nil && x.ByteLength != nil {
		return *x.ByteLength
	}
	return 0
}

func (x *CreateKeyRequest) GetOwnerId() string {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return ""
}

func (x *CreateKeyRequest) GetMeta() string {
	if x != nil && x.Meta != nil {
		return *x.Meta
	}
	return ""
}

func (x *CreateKeyRequest) GetExpires() int64 {
	if x != nil && x.Expires != nil {
		return *x.Expires
	}
	return 0
}

func (x *CreateKeyRequest) GetRatelimit() *Ratelimit {
	if x != nil {
		return x.Ratelimit
	}
	return nil
}

func (x *CreateKeyRequest) GetForWorkspaceId() string {
	if x != nil && x.ForWorkspaceId != nil {
		return *x.ForWorkspaceId
	}
	return ""
}

func (x *CreateKeyRequest) GetRemaining() int32 {
	if x != nil && x.Remaining != nil {
		return *x.Remaining
	}
	return 0
}

type CreateKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CreateKeyResponse) Reset() {
	*x = CreateKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_keys_v1_keys_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeyResponse) ProtoMessage() {}

func (x *CreateKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_keys_v1_keys_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeyResponse.ProtoReflect.Descriptor instead.
func (*CreateKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_keys_v1_keys_proto_rawDescGZIP(), []int{2}
}

func (x *CreateKeyResponse) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *CreateKeyResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	KeyAuthId      string     `protobuf:"bytes,2,opt,name=key_auth_id,json=keyAuthId,proto3" json:"key_auth_id,omitempty"`
	WorkspaceId    string     `protobuf:"bytes,3,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Name           *string    `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Hash           string     `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	Start          string     `protobuf:"bytes,6,opt,name=start,proto3" json:"start,omitempty"`
	OwnerId        *string    `protobuf:"bytes,7,opt,name=owner_id,json=ownerId,proto3,oneof" json:"owner_id,omitempty"`
	Meta           *string    `protobuf:"bytes,8,opt,name=meta,proto3,oneof" json:"meta,omitempty"`
	CreatedAt      int64      `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Expires        *int64     `protobuf:"varint,10,opt,name=expires,proto3,oneof" json:"expires,omitempty"`
	Ratelimit      *Ratelimit `protobuf:"bytes,11,opt,name=ratelimit,proto3,oneof" json:"ratelimit,omitempty"`
	ForWorkspaceId *string    `protobuf:"bytes,12,opt,name=for_workspace_id,json=forWorkspaceId,proto3,oneof" json:"for_workspace_id,omitempty"`
	Remaining      *int32     `protobuf:"varint,13,opt,name=remaining,proto3,oneof" json:"remaining,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_keys_v1_keys_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_proto_keys_v1_keys_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_proto_keys_v1_keys_proto_rawDescGZIP(), []int{3}
}

func (x *Key) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Key) GetKeyAuthId() string {
	if x != nil {
		return x.KeyAuthId
	}
	return ""
}

func (x *Key) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *Key) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Key) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Key) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Key) GetOwnerId() string {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return ""
}

func (x *Key) GetMeta() string {
	if x != nil && x.Meta != nil {
		return *x.Meta
	}
	return ""
}

func (x *Key) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Key) GetExpires() int64 {
	if x != nil && x.Expires != nil {
		return *x.Expires
	}
	return 0
}

func (x *Key) GetRatelimit() *Ratelimit {
	if x != nil {
		return x.Ratelimit
	}
	return nil
}

func (x *Key) GetForWorkspaceId() string {
	if x != nil && x.ForWorkspaceId != nil {
		return *x.ForWorkspaceId
	}
	return ""
}

func (x *Key) GetRemaining() int32 {
	if x != nil && x.Remaining != nil {
		return *x.Remaining
	}
	return 0
}

var File_proto_keys_v1_keys_proto protoreflect.FileDescriptor

var file_proto_keys_v1_keys_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x28, 0x01, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a,
	0x02, 0x28, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x30, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x28,
	0x01, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x22, 0xaa, 0x04, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x09, 0x6b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x2d, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xba, 0x48, 0x04, 0x1a, 0x02, 0x28, 0x10, 0x48,
	0x02, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x88, 0x01, 0x01,
	0x12, 0x1e, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x07, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x09, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x06, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x66, 0x6f, 0x72, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x07, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x48, 0x08, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x5f,
	0x66, 0x6f, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x3c,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x98, 0x04, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a,
	0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x6b, 0x65, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03,
	0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x09,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x04, 0x52, 0x09, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x66, 0x6f, 0x72,
	0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x09, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2a, 0x67, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x41, 0x54, 0x45,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x41, 0x54, 0x45,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x41, 0x53, 0x54, 0x10,
	0x01, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x41, 0x54, 0x45, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x02,
	0x32, 0x5f, 0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x50, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x75, 0x6e, 0x6b, 0x65, 0x79, 0x65, 0x64, 0x2f, 0x75, 0x6e, 0x6b, 0x65, 0x79, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x65, 0x79, 0x73, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_keys_v1_keys_proto_rawDescOnce sync.Once
	file_proto_keys_v1_keys_proto_rawDescData = file_proto_keys_v1_keys_proto_rawDesc
)

func file_proto_keys_v1_keys_proto_rawDescGZIP() []byte {
	file_proto_keys_v1_keys_proto_rawDescOnce.Do(func() {
		file_proto_keys_v1_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_keys_v1_keys_proto_rawDescData)
	})
	return file_proto_keys_v1_keys_proto_rawDescData
}

var file_proto_keys_v1_keys_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_keys_v1_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_keys_v1_keys_proto_goTypes = []interface{}{
	(RatelimitType)(0),        // 0: proto.keys.v1.RatelimitType
	(*Ratelimit)(nil),         // 1: proto.keys.v1.Ratelimit
	(*CreateKeyRequest)(nil),  // 2: proto.keys.v1.CreateKeyRequest
	(*CreateKeyResponse)(nil), // 3: proto.keys.v1.CreateKeyResponse
	(*Key)(nil),               // 4: proto.keys.v1.Key
}
var file_proto_keys_v1_keys_proto_depIdxs = []int32{
	0, // 0: proto.keys.v1.Ratelimit.type:type_name -> proto.keys.v1.RatelimitType
	1, // 1: proto.keys.v1.CreateKeyRequest.ratelimit:type_name -> proto.keys.v1.Ratelimit
	1, // 2: proto.keys.v1.Key.ratelimit:type_name -> proto.keys.v1.Ratelimit
	2, // 3: proto.keys.v1.KeysService.CreateKey:input_type -> proto.keys.v1.CreateKeyRequest
	3, // 4: proto.keys.v1.KeysService.CreateKey:output_type -> proto.keys.v1.CreateKeyResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_keys_v1_keys_proto_init() }
func file_proto_keys_v1_keys_proto_init() {
	if File_proto_keys_v1_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_keys_v1_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ratelimit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_keys_v1_keys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_keys_v1_keys_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_keys_v1_keys_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_keys_v1_keys_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_proto_keys_v1_keys_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_keys_v1_keys_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_keys_v1_keys_proto_goTypes,
		DependencyIndexes: file_proto_keys_v1_keys_proto_depIdxs,
		EnumInfos:         file_proto_keys_v1_keys_proto_enumTypes,
		MessageInfos:      file_proto_keys_v1_keys_proto_msgTypes,
	}.Build()
	File_proto_keys_v1_keys_proto = out.File
	file_proto_keys_v1_keys_proto_rawDesc = nil
	file_proto_keys_v1_keys_proto_goTypes = nil
	file_proto_keys_v1_keys_proto_depIdxs = nil
}
