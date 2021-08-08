// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: models/model_ns_node.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// 域名服务节点
type NSNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsOn        bool       `protobuf:"varint,3,opt,name=isOn,proto3" json:"isOn,omitempty"`
	UniqueId    string     `protobuf:"bytes,4,opt,name=uniqueId,proto3" json:"uniqueId,omitempty"`
	Secret      string     `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
	StatusJSON  []byte     `protobuf:"bytes,6,opt,name=statusJSON,proto3" json:"statusJSON,omitempty"`
	IsInstalled bool       `protobuf:"varint,7,opt,name=isInstalled,proto3" json:"isInstalled,omitempty"`
	InstallDir  string     `protobuf:"bytes,9,opt,name=installDir,proto3" json:"installDir,omitempty"`
	IsUp        bool       `protobuf:"varint,8,opt,name=isUp,proto3" json:"isUp,omitempty"`
	IsActive    bool       `protobuf:"varint,10,opt,name=isActive,proto3" json:"isActive,omitempty"`
	NsCluster   *NSCluster `protobuf:"bytes,32,opt,name=nsCluster,proto3" json:"nsCluster,omitempty"`
	//NodeLogin login = 33;
	InstallStatus *NodeInstallStatus `protobuf:"bytes,34,opt,name=installStatus,proto3" json:"installStatus,omitempty"`
}

func (x *NSNode) Reset() {
	*x = NSNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ns_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSNode) ProtoMessage() {}

func (x *NSNode) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ns_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSNode.ProtoReflect.Descriptor instead.
func (*NSNode) Descriptor() ([]byte, []int) {
	return file_models_model_ns_node_proto_rawDescGZIP(), []int{0}
}

func (x *NSNode) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NSNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NSNode) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *NSNode) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *NSNode) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *NSNode) GetStatusJSON() []byte {
	if x != nil {
		return x.StatusJSON
	}
	return nil
}

func (x *NSNode) GetIsInstalled() bool {
	if x != nil {
		return x.IsInstalled
	}
	return false
}

func (x *NSNode) GetInstallDir() string {
	if x != nil {
		return x.InstallDir
	}
	return ""
}

func (x *NSNode) GetIsUp() bool {
	if x != nil {
		return x.IsUp
	}
	return false
}

func (x *NSNode) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *NSNode) GetNsCluster() *NSCluster {
	if x != nil {
		return x.NsCluster
	}
	return nil
}

func (x *NSNode) GetInstallStatus() *NodeInstallStatus {
	if x != nil {
		return x.InstallStatus
	}
	return nil
}

var File_models_model_ns_node_proto protoreflect.FileDescriptor

var file_models_model_ns_node_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x02, 0x0a, 0x06, 0x4e, 0x53, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x20,
	0x0a, 0x0b, 0x69, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x44, 0x69, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x44, 0x69, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x55, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x69, 0x73, 0x55, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x2b, 0x0a, 0x09, 0x6e, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x09, 0x6e, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a,
	0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x22,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ns_node_proto_rawDescOnce sync.Once
	file_models_model_ns_node_proto_rawDescData = file_models_model_ns_node_proto_rawDesc
)

func file_models_model_ns_node_proto_rawDescGZIP() []byte {
	file_models_model_ns_node_proto_rawDescOnce.Do(func() {
		file_models_model_ns_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ns_node_proto_rawDescData)
	})
	return file_models_model_ns_node_proto_rawDescData
}

var file_models_model_ns_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ns_node_proto_goTypes = []interface{}{
	(*NSNode)(nil),            // 0: pb.NSNode
	(*NSCluster)(nil),         // 1: pb.NSCluster
	(*NodeInstallStatus)(nil), // 2: pb.NodeInstallStatus
}
var file_models_model_ns_node_proto_depIdxs = []int32{
	1, // 0: pb.NSNode.nsCluster:type_name -> pb.NSCluster
	2, // 1: pb.NSNode.installStatus:type_name -> pb.NodeInstallStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_model_ns_node_proto_init() }
func file_models_model_ns_node_proto_init() {
	if File_models_model_ns_node_proto != nil {
		return
	}
	file_models_model_ns_cluster_proto_init()
	file_models_model_node_install_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_ns_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSNode); i {
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
			RawDescriptor: file_models_model_ns_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ns_node_proto_goTypes,
		DependencyIndexes: file_models_model_ns_node_proto_depIdxs,
		MessageInfos:      file_models_model_ns_node_proto_msgTypes,
	}.Build()
	File_models_model_ns_node_proto = out.File
	file_models_model_ns_node_proto_rawDesc = nil
	file_models_model_ns_node_proto_goTypes = nil
	file_models_model_ns_node_proto_depIdxs = nil
}
