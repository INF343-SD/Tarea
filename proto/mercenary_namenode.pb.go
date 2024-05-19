// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0--rc1
// source: proto/mercenary_namenode.proto

package __

import (
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

type MercenaryTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codigo   string `protobuf:"bytes,1,opt,name=codigo,proto3" json:"codigo,omitempty"`
	Piso     string `protobuf:"bytes,2,opt,name=piso,proto3" json:"piso,omitempty"`
	Decision string `protobuf:"bytes,3,opt,name=decision,proto3" json:"decision,omitempty"`
}

func (x *MercenaryTask) Reset() {
	*x = MercenaryTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mercenary_namenode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercenaryTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercenaryTask) ProtoMessage() {}

func (x *MercenaryTask) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mercenary_namenode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercenaryTask.ProtoReflect.Descriptor instead.
func (*MercenaryTask) Descriptor() ([]byte, []int) {
	return file_proto_mercenary_namenode_proto_rawDescGZIP(), []int{0}
}

func (x *MercenaryTask) GetCodigo() string {
	if x != nil {
		return x.Codigo
	}
	return ""
}

func (x *MercenaryTask) GetPiso() string {
	if x != nil {
		return x.Piso
	}
	return ""
}

func (x *MercenaryTask) GetDecision() string {
	if x != nil {
		return x.Decision
	}
	return ""
}

type MercenaryTaskAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MercenaryTaskAck) Reset() {
	*x = MercenaryTaskAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mercenary_namenode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercenaryTaskAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercenaryTaskAck) ProtoMessage() {}

func (x *MercenaryTaskAck) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mercenary_namenode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercenaryTaskAck.ProtoReflect.Descriptor instead.
func (*MercenaryTaskAck) Descriptor() ([]byte, []int) {
	return file_proto_mercenary_namenode_proto_rawDescGZIP(), []int{1}
}

func (x *MercenaryTaskAck) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_mercenary_namenode_proto protoreflect.FileDescriptor

var file_proto_mercenary_namenode_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x57, 0x0a, 0x0d, 0x4d, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x64,
	0x69, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x73, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x69, 0x73, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x10, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x79,
	0x54, 0x61, 0x73, 0x6b, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x59, 0x0a, 0x14, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x79, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x53, 0x65, 0x6e,
	0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x1a,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01,
	0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mercenary_namenode_proto_rawDescOnce sync.Once
	file_proto_mercenary_namenode_proto_rawDescData = file_proto_mercenary_namenode_proto_rawDesc
)

func file_proto_mercenary_namenode_proto_rawDescGZIP() []byte {
	file_proto_mercenary_namenode_proto_rawDescOnce.Do(func() {
		file_proto_mercenary_namenode_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mercenary_namenode_proto_rawDescData)
	})
	return file_proto_mercenary_namenode_proto_rawDescData
}

var file_proto_mercenary_namenode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_mercenary_namenode_proto_goTypes = []interface{}{
	(*MercenaryTask)(nil),    // 0: namenode.MercenaryTask
	(*MercenaryTaskAck)(nil), // 1: namenode.MercenaryTaskAck
}
var file_proto_mercenary_namenode_proto_depIdxs = []int32{
	0, // 0: namenode.MercenaryTaskService.SendTask:input_type -> namenode.MercenaryTask
	1, // 1: namenode.MercenaryTaskService.SendTask:output_type -> namenode.MercenaryTaskAck
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_mercenary_namenode_proto_init() }
func file_proto_mercenary_namenode_proto_init() {
	if File_proto_mercenary_namenode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mercenary_namenode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercenaryTask); i {
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
		file_proto_mercenary_namenode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercenaryTaskAck); i {
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
			RawDescriptor: file_proto_mercenary_namenode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mercenary_namenode_proto_goTypes,
		DependencyIndexes: file_proto_mercenary_namenode_proto_depIdxs,
		MessageInfos:      file_proto_mercenary_namenode_proto_msgTypes,
	}.Build()
	File_proto_mercenary_namenode_proto = out.File
	file_proto_mercenary_namenode_proto_rawDesc = nil
	file_proto_mercenary_namenode_proto_goTypes = nil
	file_proto_mercenary_namenode_proto_depIdxs = nil
}
