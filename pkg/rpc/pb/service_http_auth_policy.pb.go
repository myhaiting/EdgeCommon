// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_http_auth_policy.proto

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

// 创建策略
type CreateHTTPAuthPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ParamsJSON []byte `protobuf:"bytes,3,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
}

func (x *CreateHTTPAuthPolicyRequest) Reset() {
	*x = CreateHTTPAuthPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_auth_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPAuthPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPAuthPolicyRequest) ProtoMessage() {}

func (x *CreateHTTPAuthPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_auth_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPAuthPolicyRequest.ProtoReflect.Descriptor instead.
func (*CreateHTTPAuthPolicyRequest) Descriptor() ([]byte, []int) {
	return file_service_http_auth_policy_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHTTPAuthPolicyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateHTTPAuthPolicyRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateHTTPAuthPolicyRequest) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

type CreateHTTPAuthPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpAuthPolicyId int64 `protobuf:"varint,1,opt,name=httpAuthPolicyId,proto3" json:"httpAuthPolicyId,omitempty"`
}

func (x *CreateHTTPAuthPolicyResponse) Reset() {
	*x = CreateHTTPAuthPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_auth_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPAuthPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPAuthPolicyResponse) ProtoMessage() {}

func (x *CreateHTTPAuthPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_auth_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPAuthPolicyResponse.ProtoReflect.Descriptor instead.
func (*CreateHTTPAuthPolicyResponse) Descriptor() ([]byte, []int) {
	return file_service_http_auth_policy_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHTTPAuthPolicyResponse) GetHttpAuthPolicyId() int64 {
	if x != nil {
		return x.HttpAuthPolicyId
	}
	return 0
}

// 修改策略
type UpdateHTTPAuthPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpAuthPolicyId int64  `protobuf:"varint,1,opt,name=httpAuthPolicyId,proto3" json:"httpAuthPolicyId,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParamsJSON       []byte `protobuf:"bytes,3,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
	IsOn             bool   `protobuf:"varint,4,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateHTTPAuthPolicyRequest) Reset() {
	*x = UpdateHTTPAuthPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_auth_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPAuthPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPAuthPolicyRequest) ProtoMessage() {}

func (x *UpdateHTTPAuthPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_auth_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPAuthPolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPAuthPolicyRequest) Descriptor() ([]byte, []int) {
	return file_service_http_auth_policy_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHTTPAuthPolicyRequest) GetHttpAuthPolicyId() int64 {
	if x != nil {
		return x.HttpAuthPolicyId
	}
	return 0
}

func (x *UpdateHTTPAuthPolicyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateHTTPAuthPolicyRequest) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

func (x *UpdateHTTPAuthPolicyRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 查找策略信息
type FindEnabledHTTPAuthPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpAuthPolicyId int64 `protobuf:"varint,1,opt,name=httpAuthPolicyId,proto3" json:"httpAuthPolicyId,omitempty"`
}

func (x *FindEnabledHTTPAuthPolicyRequest) Reset() {
	*x = FindEnabledHTTPAuthPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_auth_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPAuthPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPAuthPolicyRequest) ProtoMessage() {}

func (x *FindEnabledHTTPAuthPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_auth_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPAuthPolicyRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPAuthPolicyRequest) Descriptor() ([]byte, []int) {
	return file_service_http_auth_policy_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledHTTPAuthPolicyRequest) GetHttpAuthPolicyId() int64 {
	if x != nil {
		return x.HttpAuthPolicyId
	}
	return 0
}

type FindEnabledHTTPAuthPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpAuthPolicy *HTTPAuthPolicy `protobuf:"bytes,1,opt,name=httpAuthPolicy,proto3" json:"httpAuthPolicy,omitempty"`
}

func (x *FindEnabledHTTPAuthPolicyResponse) Reset() {
	*x = FindEnabledHTTPAuthPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_auth_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPAuthPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPAuthPolicyResponse) ProtoMessage() {}

func (x *FindEnabledHTTPAuthPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_auth_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPAuthPolicyResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPAuthPolicyResponse) Descriptor() ([]byte, []int) {
	return file_service_http_auth_policy_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledHTTPAuthPolicyResponse) GetHttpAuthPolicy() *HTTPAuthPolicy {
	if x != nil {
		return x.HttpAuthPolicy
	}
	return nil
}

var File_service_http_auth_policy_proto protoreflect.FileDescriptor

var file_service_http_auth_policy_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x23, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54,
	0x54, 0x50, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x4a, 0x0a, 0x1c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x68,
	0x74, 0x74, 0x70, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x68, 0x74, 0x74, 0x70, 0x41, 0x75, 0x74, 0x68, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x68, 0x74, 0x74, 0x70, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x68, 0x74, 0x74, 0x70, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x22, 0x4e, 0x0a, 0x20, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x41, 0x75,
	0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x10, 0x68, 0x74, 0x74, 0x70, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x68, 0x74, 0x74, 0x70, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x21, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x41, 0x75,
	0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x0e, 0x68, 0x74, 0x74, 0x70, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54,
	0x54, 0x50, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x68, 0x74,
	0x74, 0x70, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x32, 0xa5, 0x02, 0x0a,
	0x15, 0x48, 0x54, 0x54, 0x50, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x48, 0x54, 0x54, 0x50, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1f,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x41, 0x75,
	0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x68, 0x0a, 0x19, 0x66, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54,
	0x54, 0x50, 0x41, 0x75, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_http_auth_policy_proto_rawDescOnce sync.Once
	file_service_http_auth_policy_proto_rawDescData = file_service_http_auth_policy_proto_rawDesc
)

func file_service_http_auth_policy_proto_rawDescGZIP() []byte {
	file_service_http_auth_policy_proto_rawDescOnce.Do(func() {
		file_service_http_auth_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_http_auth_policy_proto_rawDescData)
	})
	return file_service_http_auth_policy_proto_rawDescData
}

var file_service_http_auth_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_http_auth_policy_proto_goTypes = []interface{}{
	(*CreateHTTPAuthPolicyRequest)(nil),       // 0: pb.CreateHTTPAuthPolicyRequest
	(*CreateHTTPAuthPolicyResponse)(nil),      // 1: pb.CreateHTTPAuthPolicyResponse
	(*UpdateHTTPAuthPolicyRequest)(nil),       // 2: pb.UpdateHTTPAuthPolicyRequest
	(*FindEnabledHTTPAuthPolicyRequest)(nil),  // 3: pb.FindEnabledHTTPAuthPolicyRequest
	(*FindEnabledHTTPAuthPolicyResponse)(nil), // 4: pb.FindEnabledHTTPAuthPolicyResponse
	(*HTTPAuthPolicy)(nil),                    // 5: pb.HTTPAuthPolicy
	(*RPCSuccess)(nil),                        // 6: pb.RPCSuccess
}
var file_service_http_auth_policy_proto_depIdxs = []int32{
	5, // 0: pb.FindEnabledHTTPAuthPolicyResponse.httpAuthPolicy:type_name -> pb.HTTPAuthPolicy
	0, // 1: pb.HTTPAuthPolicyService.createHTTPAuthPolicy:input_type -> pb.CreateHTTPAuthPolicyRequest
	2, // 2: pb.HTTPAuthPolicyService.updateHTTPAuthPolicy:input_type -> pb.UpdateHTTPAuthPolicyRequest
	3, // 3: pb.HTTPAuthPolicyService.findEnabledHTTPAuthPolicy:input_type -> pb.FindEnabledHTTPAuthPolicyRequest
	1, // 4: pb.HTTPAuthPolicyService.createHTTPAuthPolicy:output_type -> pb.CreateHTTPAuthPolicyResponse
	6, // 5: pb.HTTPAuthPolicyService.updateHTTPAuthPolicy:output_type -> pb.RPCSuccess
	4, // 6: pb.HTTPAuthPolicyService.findEnabledHTTPAuthPolicy:output_type -> pb.FindEnabledHTTPAuthPolicyResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_http_auth_policy_proto_init() }
func file_service_http_auth_policy_proto_init() {
	if File_service_http_auth_policy_proto != nil {
		return
	}
	file_models_model_http_auth_policy_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_http_auth_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPAuthPolicyRequest); i {
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
		file_service_http_auth_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPAuthPolicyResponse); i {
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
		file_service_http_auth_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPAuthPolicyRequest); i {
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
		file_service_http_auth_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPAuthPolicyRequest); i {
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
		file_service_http_auth_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPAuthPolicyResponse); i {
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
			RawDescriptor: file_service_http_auth_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_http_auth_policy_proto_goTypes,
		DependencyIndexes: file_service_http_auth_policy_proto_depIdxs,
		MessageInfos:      file_service_http_auth_policy_proto_msgTypes,
	}.Build()
	File_service_http_auth_policy_proto = out.File
	file_service_http_auth_policy_proto_rawDesc = nil
	file_service_http_auth_policy_proto_goTypes = nil
	file_service_http_auth_policy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTTPAuthPolicyServiceClient is the client API for HTTPAuthPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPAuthPolicyServiceClient interface {
	// 创建策略
	CreateHTTPAuthPolicy(ctx context.Context, in *CreateHTTPAuthPolicyRequest, opts ...grpc.CallOption) (*CreateHTTPAuthPolicyResponse, error)
	// 修改策略
	UpdateHTTPAuthPolicy(ctx context.Context, in *UpdateHTTPAuthPolicyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找策略信息
	FindEnabledHTTPAuthPolicy(ctx context.Context, in *FindEnabledHTTPAuthPolicyRequest, opts ...grpc.CallOption) (*FindEnabledHTTPAuthPolicyResponse, error)
}

type hTTPAuthPolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPAuthPolicyServiceClient(cc grpc.ClientConnInterface) HTTPAuthPolicyServiceClient {
	return &hTTPAuthPolicyServiceClient{cc}
}

func (c *hTTPAuthPolicyServiceClient) CreateHTTPAuthPolicy(ctx context.Context, in *CreateHTTPAuthPolicyRequest, opts ...grpc.CallOption) (*CreateHTTPAuthPolicyResponse, error) {
	out := new(CreateHTTPAuthPolicyResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPAuthPolicyService/createHTTPAuthPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPAuthPolicyServiceClient) UpdateHTTPAuthPolicy(ctx context.Context, in *UpdateHTTPAuthPolicyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPAuthPolicyService/updateHTTPAuthPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPAuthPolicyServiceClient) FindEnabledHTTPAuthPolicy(ctx context.Context, in *FindEnabledHTTPAuthPolicyRequest, opts ...grpc.CallOption) (*FindEnabledHTTPAuthPolicyResponse, error) {
	out := new(FindEnabledHTTPAuthPolicyResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPAuthPolicyService/findEnabledHTTPAuthPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPAuthPolicyServiceServer is the server API for HTTPAuthPolicyService service.
type HTTPAuthPolicyServiceServer interface {
	// 创建策略
	CreateHTTPAuthPolicy(context.Context, *CreateHTTPAuthPolicyRequest) (*CreateHTTPAuthPolicyResponse, error)
	// 修改策略
	UpdateHTTPAuthPolicy(context.Context, *UpdateHTTPAuthPolicyRequest) (*RPCSuccess, error)
	// 查找策略信息
	FindEnabledHTTPAuthPolicy(context.Context, *FindEnabledHTTPAuthPolicyRequest) (*FindEnabledHTTPAuthPolicyResponse, error)
}

// UnimplementedHTTPAuthPolicyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHTTPAuthPolicyServiceServer struct {
}

func (*UnimplementedHTTPAuthPolicyServiceServer) CreateHTTPAuthPolicy(context.Context, *CreateHTTPAuthPolicyRequest) (*CreateHTTPAuthPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPAuthPolicy not implemented")
}
func (*UnimplementedHTTPAuthPolicyServiceServer) UpdateHTTPAuthPolicy(context.Context, *UpdateHTTPAuthPolicyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPAuthPolicy not implemented")
}
func (*UnimplementedHTTPAuthPolicyServiceServer) FindEnabledHTTPAuthPolicy(context.Context, *FindEnabledHTTPAuthPolicyRequest) (*FindEnabledHTTPAuthPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPAuthPolicy not implemented")
}

func RegisterHTTPAuthPolicyServiceServer(s *grpc.Server, srv HTTPAuthPolicyServiceServer) {
	s.RegisterService(&_HTTPAuthPolicyService_serviceDesc, srv)
}

func _HTTPAuthPolicyService_CreateHTTPAuthPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPAuthPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAuthPolicyServiceServer).CreateHTTPAuthPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPAuthPolicyService/CreateHTTPAuthPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAuthPolicyServiceServer).CreateHTTPAuthPolicy(ctx, req.(*CreateHTTPAuthPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPAuthPolicyService_UpdateHTTPAuthPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPAuthPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAuthPolicyServiceServer).UpdateHTTPAuthPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPAuthPolicyService/UpdateHTTPAuthPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAuthPolicyServiceServer).UpdateHTTPAuthPolicy(ctx, req.(*UpdateHTTPAuthPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPAuthPolicyService_FindEnabledHTTPAuthPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPAuthPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAuthPolicyServiceServer).FindEnabledHTTPAuthPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPAuthPolicyService/FindEnabledHTTPAuthPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAuthPolicyServiceServer).FindEnabledHTTPAuthPolicy(ctx, req.(*FindEnabledHTTPAuthPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTPAuthPolicyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPAuthPolicyService",
	HandlerType: (*HTTPAuthPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHTTPAuthPolicy",
			Handler:    _HTTPAuthPolicyService_CreateHTTPAuthPolicy_Handler,
		},
		{
			MethodName: "updateHTTPAuthPolicy",
			Handler:    _HTTPAuthPolicyService_UpdateHTTPAuthPolicy_Handler,
		},
		{
			MethodName: "findEnabledHTTPAuthPolicy",
			Handler:    _HTTPAuthPolicyService_FindEnabledHTTPAuthPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_auth_policy.proto",
}
