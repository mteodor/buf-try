// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: entity/agents.proto

package agents

import (
	_ "github.com/gogo/protobuf/gogoproto"
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

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_agents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_entity_agents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_entity_agents_proto_rawDescGZIP(), []int{0}
}

func (x *Entity) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

var File_entity_agents_proto protoreflect.FileDescriptor

var file_entity_agents_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3f, 0x0a,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1f, 0xf2, 0xde, 0x1f, 0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x76,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x0b, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x74, 0x65, 0x6f, 0x64, 0x6f, 0x72, 0x6f,
	0x76, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x72, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73,
	0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0xca,
	0x02, 0x06, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0xe2, 0x02, 0x12, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_agents_proto_rawDescOnce sync.Once
	file_entity_agents_proto_rawDescData = file_entity_agents_proto_rawDesc
)

func file_entity_agents_proto_rawDescGZIP() []byte {
	file_entity_agents_proto_rawDescOnce.Do(func() {
		file_entity_agents_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_agents_proto_rawDescData)
	})
	return file_entity_agents_proto_rawDescData
}

var file_entity_agents_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entity_agents_proto_goTypes = []interface{}{
	(*Entity)(nil), // 0: agents.Entity
}
var file_entity_agents_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_entity_agents_proto_init() }
func file_entity_agents_proto_init() {
	if File_entity_agents_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_agents_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
			RawDescriptor: file_entity_agents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_agents_proto_goTypes,
		DependencyIndexes: file_entity_agents_proto_depIdxs,
		MessageInfos:      file_entity_agents_proto_msgTypes,
	}.Build()
	File_entity_agents_proto = out.File
	file_entity_agents_proto_rawDesc = nil
	file_entity_agents_proto_goTypes = nil
	file_entity_agents_proto_depIdxs = nil
}
