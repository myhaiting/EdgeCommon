// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_origin.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 创建源站
type CreateOriginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Addr        *NetworkAddress `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Description string          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Weight      int32           `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
	IsOn        bool            `protobuf:"varint,5,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *CreateOriginRequest) Reset() {
	*x = CreateOriginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_origin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOriginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOriginRequest) ProtoMessage() {}

func (x *CreateOriginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_origin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOriginRequest.ProtoReflect.Descriptor instead.
func (*CreateOriginRequest) Descriptor() ([]byte, []int) {
	return file_service_origin_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOriginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateOriginRequest) GetAddr() *NetworkAddress {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *CreateOriginRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateOriginRequest) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *CreateOriginRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

type CreateOriginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginId int64 `protobuf:"varint,1,opt,name=originId,proto3" json:"originId,omitempty"`
}

func (x *CreateOriginResponse) Reset() {
	*x = CreateOriginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_origin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOriginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOriginResponse) ProtoMessage() {}

func (x *CreateOriginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_origin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOriginResponse.ProtoReflect.Descriptor instead.
func (*CreateOriginResponse) Descriptor() ([]byte, []int) {
	return file_service_origin_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOriginResponse) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

// 修改源站
type UpdateOriginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginId    int64           `protobuf:"varint,1,opt,name=originId,proto3" json:"originId,omitempty"`
	Name        string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Addr        *NetworkAddress `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	Description string          `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Weight      int32           `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	IsOn        bool            `protobuf:"varint,6,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateOriginRequest) Reset() {
	*x = UpdateOriginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_origin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOriginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOriginRequest) ProtoMessage() {}

func (x *UpdateOriginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_origin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOriginRequest.ProtoReflect.Descriptor instead.
func (*UpdateOriginRequest) Descriptor() ([]byte, []int) {
	return file_service_origin_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateOriginRequest) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

func (x *UpdateOriginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateOriginRequest) GetAddr() *NetworkAddress {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *UpdateOriginRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateOriginRequest) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *UpdateOriginRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 查找单个源站信息
type FindEnabledOriginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginId int64 `protobuf:"varint,1,opt,name=originId,proto3" json:"originId,omitempty"`
}

func (x *FindEnabledOriginRequest) Reset() {
	*x = FindEnabledOriginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_origin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledOriginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledOriginRequest) ProtoMessage() {}

func (x *FindEnabledOriginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_origin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledOriginRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledOriginRequest) Descriptor() ([]byte, []int) {
	return file_service_origin_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledOriginRequest) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

type FindEnabledOriginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin *Origin `protobuf:"bytes,1,opt,name=Origin,proto3" json:"Origin,omitempty"`
}

func (x *FindEnabledOriginResponse) Reset() {
	*x = FindEnabledOriginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_origin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledOriginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledOriginResponse) ProtoMessage() {}

func (x *FindEnabledOriginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_origin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledOriginResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledOriginResponse) Descriptor() ([]byte, []int) {
	return file_service_origin_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledOriginResponse) GetOrigin() *Origin {
	if x != nil {
		return x.Origin
	}
	return nil
}

// 查找源站配置
type FindEnabledOriginConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginId int64 `protobuf:"varint,1,opt,name=originId,proto3" json:"originId,omitempty"`
}

func (x *FindEnabledOriginConfigRequest) Reset() {
	*x = FindEnabledOriginConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_origin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledOriginConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledOriginConfigRequest) ProtoMessage() {}

func (x *FindEnabledOriginConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_origin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledOriginConfigRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledOriginConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_origin_proto_rawDescGZIP(), []int{5}
}

func (x *FindEnabledOriginConfigRequest) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

type FindEnabledOriginConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginJSON []byte `protobuf:"bytes,1,opt,name=originJSON,proto3" json:"originJSON,omitempty"`
}

func (x *FindEnabledOriginConfigResponse) Reset() {
	*x = FindEnabledOriginConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_origin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledOriginConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledOriginConfigResponse) ProtoMessage() {}

func (x *FindEnabledOriginConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_origin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledOriginConfigResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledOriginConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_origin_proto_rawDescGZIP(), []int{6}
}

func (x *FindEnabledOriginConfigResponse) GetOriginJSON() []byte {
	if x != nil {
		return x.OriginJSON
	}
	return nil
}

var File_service_origin_proto protoreflect.FileDescriptor

var file_service_origin_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x70, 0x63,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9f, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x22, 0x32, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xbb, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x22, 0x36, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x19, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x22, 0x3c, 0x0a, 0x1e,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x1f, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4a, 0x53, 0x4f, 0x4e, 0x32, 0xc1, 0x02,
	0x0a, 0x0d, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x41, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12,
	0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x50, 0x0a, 0x11, 0x66,
	0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a,
	0x17, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_origin_proto_rawDescOnce sync.Once
	file_service_origin_proto_rawDescData = file_service_origin_proto_rawDesc
)

func file_service_origin_proto_rawDescGZIP() []byte {
	file_service_origin_proto_rawDescOnce.Do(func() {
		file_service_origin_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_origin_proto_rawDescData)
	})
	return file_service_origin_proto_rawDescData
}

var file_service_origin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_origin_proto_goTypes = []interface{}{
	(*CreateOriginRequest)(nil),             // 0: pb.CreateOriginRequest
	(*CreateOriginResponse)(nil),            // 1: pb.CreateOriginResponse
	(*UpdateOriginRequest)(nil),             // 2: pb.UpdateOriginRequest
	(*FindEnabledOriginRequest)(nil),        // 3: pb.FindEnabledOriginRequest
	(*FindEnabledOriginResponse)(nil),       // 4: pb.FindEnabledOriginResponse
	(*FindEnabledOriginConfigRequest)(nil),  // 5: pb.FindEnabledOriginConfigRequest
	(*FindEnabledOriginConfigResponse)(nil), // 6: pb.FindEnabledOriginConfigResponse
	(*NetworkAddress)(nil),                  // 7: pb.NetworkAddress
	(*Origin)(nil),                          // 8: pb.Origin
	(*RPCSuccess)(nil),                      // 9: pb.RPCSuccess
}
var file_service_origin_proto_depIdxs = []int32{
	7, // 0: pb.CreateOriginRequest.addr:type_name -> pb.NetworkAddress
	7, // 1: pb.UpdateOriginRequest.addr:type_name -> pb.NetworkAddress
	8, // 2: pb.FindEnabledOriginResponse.Origin:type_name -> pb.Origin
	0, // 3: pb.OriginService.createOrigin:input_type -> pb.CreateOriginRequest
	2, // 4: pb.OriginService.updateOrigin:input_type -> pb.UpdateOriginRequest
	3, // 5: pb.OriginService.findEnabledOrigin:input_type -> pb.FindEnabledOriginRequest
	5, // 6: pb.OriginService.findEnabledOriginConfig:input_type -> pb.FindEnabledOriginConfigRequest
	1, // 7: pb.OriginService.createOrigin:output_type -> pb.CreateOriginResponse
	9, // 8: pb.OriginService.updateOrigin:output_type -> pb.RPCSuccess
	4, // 9: pb.OriginService.findEnabledOrigin:output_type -> pb.FindEnabledOriginResponse
	6, // 10: pb.OriginService.findEnabledOriginConfig:output_type -> pb.FindEnabledOriginConfigResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_origin_proto_init() }
func file_service_origin_proto_init() {
	if File_service_origin_proto != nil {
		return
	}
	file_model_origin_proto_init()
	file_model_network_address_proto_init()
	file_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_origin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOriginRequest); i {
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
		file_service_origin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOriginResponse); i {
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
		file_service_origin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOriginRequest); i {
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
		file_service_origin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledOriginRequest); i {
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
		file_service_origin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledOriginResponse); i {
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
		file_service_origin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledOriginConfigRequest); i {
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
		file_service_origin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledOriginConfigResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_origin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_origin_proto_goTypes,
		DependencyIndexes: file_service_origin_proto_depIdxs,
		MessageInfos:      file_service_origin_proto_msgTypes,
	}.Build()
	File_service_origin_proto = out.File
	file_service_origin_proto_rawDesc = nil
	file_service_origin_proto_goTypes = nil
	file_service_origin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OriginServiceClient is the client API for OriginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OriginServiceClient interface {
	// 创建源站
	CreateOrigin(ctx context.Context, in *CreateOriginRequest, opts ...grpc.CallOption) (*CreateOriginResponse, error)
	// 修改源站
	UpdateOrigin(ctx context.Context, in *UpdateOriginRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个源站信息
	FindEnabledOrigin(ctx context.Context, in *FindEnabledOriginRequest, opts ...grpc.CallOption) (*FindEnabledOriginResponse, error)
	// 查找源站配置
	FindEnabledOriginConfig(ctx context.Context, in *FindEnabledOriginConfigRequest, opts ...grpc.CallOption) (*FindEnabledOriginConfigResponse, error)
}

type originServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOriginServiceClient(cc grpc.ClientConnInterface) OriginServiceClient {
	return &originServiceClient{cc}
}

func (c *originServiceClient) CreateOrigin(ctx context.Context, in *CreateOriginRequest, opts ...grpc.CallOption) (*CreateOriginResponse, error) {
	out := new(CreateOriginResponse)
	err := c.cc.Invoke(ctx, "/pb.OriginService/createOrigin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originServiceClient) UpdateOrigin(ctx context.Context, in *UpdateOriginRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.OriginService/updateOrigin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originServiceClient) FindEnabledOrigin(ctx context.Context, in *FindEnabledOriginRequest, opts ...grpc.CallOption) (*FindEnabledOriginResponse, error) {
	out := new(FindEnabledOriginResponse)
	err := c.cc.Invoke(ctx, "/pb.OriginService/findEnabledOrigin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originServiceClient) FindEnabledOriginConfig(ctx context.Context, in *FindEnabledOriginConfigRequest, opts ...grpc.CallOption) (*FindEnabledOriginConfigResponse, error) {
	out := new(FindEnabledOriginConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.OriginService/findEnabledOriginConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OriginServiceServer is the server API for OriginService service.
type OriginServiceServer interface {
	// 创建源站
	CreateOrigin(context.Context, *CreateOriginRequest) (*CreateOriginResponse, error)
	// 修改源站
	UpdateOrigin(context.Context, *UpdateOriginRequest) (*RPCSuccess, error)
	// 查找单个源站信息
	FindEnabledOrigin(context.Context, *FindEnabledOriginRequest) (*FindEnabledOriginResponse, error)
	// 查找源站配置
	FindEnabledOriginConfig(context.Context, *FindEnabledOriginConfigRequest) (*FindEnabledOriginConfigResponse, error)
}

// UnimplementedOriginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOriginServiceServer struct {
}

func (*UnimplementedOriginServiceServer) CreateOrigin(context.Context, *CreateOriginRequest) (*CreateOriginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrigin not implemented")
}
func (*UnimplementedOriginServiceServer) UpdateOrigin(context.Context, *UpdateOriginRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrigin not implemented")
}
func (*UnimplementedOriginServiceServer) FindEnabledOrigin(context.Context, *FindEnabledOriginRequest) (*FindEnabledOriginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledOrigin not implemented")
}
func (*UnimplementedOriginServiceServer) FindEnabledOriginConfig(context.Context, *FindEnabledOriginConfigRequest) (*FindEnabledOriginConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledOriginConfig not implemented")
}

func RegisterOriginServiceServer(s *grpc.Server, srv OriginServiceServer) {
	s.RegisterService(&_OriginService_serviceDesc, srv)
}

func _OriginService_CreateOrigin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOriginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginServiceServer).CreateOrigin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OriginService/CreateOrigin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginServiceServer).CreateOrigin(ctx, req.(*CreateOriginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginService_UpdateOrigin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOriginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginServiceServer).UpdateOrigin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OriginService/UpdateOrigin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginServiceServer).UpdateOrigin(ctx, req.(*UpdateOriginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginService_FindEnabledOrigin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledOriginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginServiceServer).FindEnabledOrigin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OriginService/FindEnabledOrigin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginServiceServer).FindEnabledOrigin(ctx, req.(*FindEnabledOriginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginService_FindEnabledOriginConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledOriginConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginServiceServer).FindEnabledOriginConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OriginService/FindEnabledOriginConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginServiceServer).FindEnabledOriginConfig(ctx, req.(*FindEnabledOriginConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OriginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OriginService",
	HandlerType: (*OriginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createOrigin",
			Handler:    _OriginService_CreateOrigin_Handler,
		},
		{
			MethodName: "updateOrigin",
			Handler:    _OriginService_UpdateOrigin_Handler,
		},
		{
			MethodName: "findEnabledOrigin",
			Handler:    _OriginService_FindEnabledOrigin_Handler,
		},
		{
			MethodName: "findEnabledOriginConfig",
			Handler:    _OriginService_FindEnabledOriginConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_origin.proto",
}