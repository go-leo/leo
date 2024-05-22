// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: annotations.proto

package cqrs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Responsibility int32

const (
	Responsibility_Unknown Responsibility = 0
	Responsibility_Command Responsibility = 1
	Responsibility_Query   Responsibility = 2
)

// Enum value maps for Responsibility.
var (
	Responsibility_name = map[int32]string{
		0: "Unknown",
		1: "Command",
		2: "Query",
	}
	Responsibility_value = map[string]int32{
		"Unknown": 0,
		"Command": 1,
		"Query":   2,
	}
)

func (x Responsibility) Enum() *Responsibility {
	p := new(Responsibility)
	*p = x
	return p
}

func (x Responsibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Responsibility) Descriptor() protoreflect.EnumDescriptor {
	return file_annotations_proto_enumTypes[0].Descriptor()
}

func (Responsibility) Type() protoreflect.EnumType {
	return &file_annotations_proto_enumTypes[0]
}

func (x Responsibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Responsibility.Descriptor instead.
func (Responsibility) EnumDescriptor() ([]byte, []int) {
	return file_annotations_proto_rawDescGZIP(), []int{0}
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// package the full package name of the command or query.
	Package string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	// relative the package path of the the command or query, relative to the current proto file.
	Relative string `protobuf:"bytes,2,opt,name=relative,proto3" json:"relative,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *Package) GetRelative() string {
	if x != nil {
		return x.Relative
	}
	return ""
}

var file_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*Package)(nil),
		Field:         31718192,
		Name:          "leo.cqrs.command",
		Tag:           "bytes,31718192,opt,name=command",
		Filename:      "annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*Package)(nil),
		Field:         31718193,
		Name:          "leo.cqrs.query",
		Tag:           "bytes,31718193,opt,name=query",
		Filename:      "annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Responsibility)(nil),
		Field:         3171819,
		Name:          "leo.cqrs.responsibility",
		Tag:           "varint,3171819,opt,name=responsibility,enum=leo.cqrs.Responsibility",
		Filename:      "annotations.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// Define command package information
	//
	// optional leo.cqrs.Package command = 31718192;
	E_Command = &file_annotations_proto_extTypes[0]
	// Define query package information
	//
	// optional leo.cqrs.Package query = 31718193;
	E_Query = &file_annotations_proto_extTypes[1]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// responsibility define command or query api
	//
	// optional leo.cqrs.Responsibility responsibility = 3171819;
	E_Responsibility = &file_annotations_proto_extTypes[2]
)

var File_annotations_proto protoreflect.FileDescriptor

var file_annotations_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x65, 0x6f, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3f, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x2a, 0x35, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x10, 0x02, 0x3a, 0x4f, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xb0, 0xf6, 0x8f, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c,
	0x65, 0x6f, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x3a, 0x4b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xb1, 0xf6, 0x8f, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x65,
	0x6f, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x63, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xeb, 0xcb, 0xc1, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x46, 0x0a, 0x0c, 0x64, 0x65,
	0x76, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x42, 0x09, 0x43, 0x51, 0x52, 0x53,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x6c, 0x65, 0x6f, 0x2f, 0x76,
	0x33, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x3b, 0x63, 0x71, 0x72, 0x73, 0xa2, 0x02, 0x04, 0x43, 0x51,
	0x52, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_annotations_proto_rawDescOnce sync.Once
	file_annotations_proto_rawDescData = file_annotations_proto_rawDesc
)

func file_annotations_proto_rawDescGZIP() []byte {
	file_annotations_proto_rawDescOnce.Do(func() {
		file_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_annotations_proto_rawDescData)
	})
	return file_annotations_proto_rawDescData
}

var file_annotations_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_annotations_proto_goTypes = []interface{}{
	(Responsibility)(0),                 // 0: leo.cqrs.Responsibility
	(*Package)(nil),                     // 1: leo.cqrs.Package
	(*descriptorpb.ServiceOptions)(nil), // 2: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 3: google.protobuf.MethodOptions
}
var file_annotations_proto_depIdxs = []int32{
	2, // 0: leo.cqrs.command:extendee -> google.protobuf.ServiceOptions
	2, // 1: leo.cqrs.query:extendee -> google.protobuf.ServiceOptions
	3, // 2: leo.cqrs.responsibility:extendee -> google.protobuf.MethodOptions
	1, // 3: leo.cqrs.command:type_name -> leo.cqrs.Package
	1, // 4: leo.cqrs.query:type_name -> leo.cqrs.Package
	0, // 5: leo.cqrs.responsibility:type_name -> leo.cqrs.Responsibility
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	3, // [3:6] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_annotations_proto_init() }
func file_annotations_proto_init() {
	if File_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
			RawDescriptor: file_annotations_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_annotations_proto_goTypes,
		DependencyIndexes: file_annotations_proto_depIdxs,
		EnumInfos:         file_annotations_proto_enumTypes,
		MessageInfos:      file_annotations_proto_msgTypes,
		ExtensionInfos:    file_annotations_proto_extTypes,
	}.Build()
	File_annotations_proto = out.File
	file_annotations_proto_rawDesc = nil
	file_annotations_proto_goTypes = nil
	file_annotations_proto_depIdxs = nil
}
