// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: status/v1/error.proto

package status

import (
	_ "github.com/go-leo/leo/v3/proto/leo/status"
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

type Errors int32

const (
	Errors_InvalidName        Errors = 0
	Errors_FileDownloadFailed Errors = 1
	Errors_FileUploadFailed   Errors = 2
)

// Enum value maps for Errors.
var (
	Errors_name = map[int32]string{
		0: "InvalidName",
		1: "FileDownloadFailed",
		2: "FileUploadFailed",
	}
	Errors_value = map[string]int32{
		"InvalidName":        0,
		"FileDownloadFailed": 1,
		"FileUploadFailed":   2,
	}
)

func (x Errors) Enum() *Errors {
	p := new(Errors)
	*p = x
	return p
}

func (x Errors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Errors) Descriptor() protoreflect.EnumDescriptor {
	return file_status_v1_error_proto_enumTypes[0].Descriptor()
}

func (Errors) Type() protoreflect.EnumType {
	return &file_status_v1_error_proto_enumTypes[0]
}

func (x Errors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Errors.Descriptor instead.
func (Errors) EnumDescriptor() ([]byte, []int) {
	return file_status_v1_error_proto_rawDescGZIP(), []int{0}
}

var File_status_v1_error_proto protoreflect.FileDescriptor

var file_status_v1_error_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x1a, 0x1c, 0x6c, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x7b, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x00, 0x1a, 0x14, 0xb0, 0xb7, 0x22,
	0x03, 0xda, 0x94, 0x28, 0x0c, 0xe5, 0x90, 0x8d, 0xe7, 0xa7, 0xb0, 0xe4, 0xb8, 0xba, 0xe7, 0xa9,
	0xba, 0x12, 0x2e, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x1a, 0x16, 0xda, 0x94, 0x28, 0x12, 0xe6,
	0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe4, 0xb8, 0x8b, 0xe8, 0xbd, 0xbd, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4,
	0xa5, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x1a, 0x04, 0xa0, 0xe5, 0x1f, 0x0d, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6c,
	0x65, 0x6f, 0x2f, 0x6c, 0x65, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_status_v1_error_proto_rawDescOnce sync.Once
	file_status_v1_error_proto_rawDescData = file_status_v1_error_proto_rawDesc
)

func file_status_v1_error_proto_rawDescGZIP() []byte {
	file_status_v1_error_proto_rawDescOnce.Do(func() {
		file_status_v1_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_status_v1_error_proto_rawDescData)
	})
	return file_status_v1_error_proto_rawDescData
}

var file_status_v1_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_status_v1_error_proto_goTypes = []interface{}{
	(Errors)(0), // 0: leo.example.status.errors.Errors
}
var file_status_v1_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_status_v1_error_proto_init() }
func file_status_v1_error_proto_init() {
	if File_status_v1_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_status_v1_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_status_v1_error_proto_goTypes,
		DependencyIndexes: file_status_v1_error_proto_depIdxs,
		EnumInfos:         file_status_v1_error_proto_enumTypes,
	}.Build()
	File_status_v1_error_proto = out.File
	file_status_v1_error_proto_rawDesc = nil
	file_status_v1_error_proto_goTypes = nil
	file_status_v1_error_proto_depIdxs = nil
}
