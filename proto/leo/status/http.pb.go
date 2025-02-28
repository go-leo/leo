// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: leo/status/http.proto

package status

import (
	code "google.golang.org/genproto/googleapis/rpc/code"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// HttpBody is the http body.
// see: https://google.aip.dev/193 HTTP/1.1+JSON representation
type HttpBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *HttpBody_Status `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *HttpBody) Reset() {
	*x = HttpBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leo_status_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpBody) ProtoMessage() {}

func (x *HttpBody) ProtoReflect() protoreflect.Message {
	mi := &file_leo_status_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpBody.ProtoReflect.Descriptor instead.
func (*HttpBody) Descriptor() ([]byte, []int) {
	return file_leo_status_http_proto_rawDescGZIP(), []int{0}
}

func (x *HttpBody) GetError() *HttpBody_Status {
	if x != nil {
		return x.Error
	}
	return nil
}

type HttpBody_Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the enum version for `google.rpc.Status.code`.
	Status code.Code `protobuf:"varint,1,opt,name=status,proto3,enum=google.rpc.Code" json:"status,omitempty"`
	// This corresponds to `google.rpc.Status.message`.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The HTTP status code that corresponds to `google.rpc.Status.code`.
	Code int32 `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	// This distinguish between two Status objects as being the same when
	// both code and status are identical.
	Identifier string `protobuf:"bytes,4,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// This corresponds to `google.rpc.Status.details`.
	Details []*anypb.Any `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *HttpBody_Status) Reset() {
	*x = HttpBody_Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leo_status_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpBody_Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpBody_Status) ProtoMessage() {}

func (x *HttpBody_Status) ProtoReflect() protoreflect.Message {
	mi := &file_leo_status_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpBody_Status.ProtoReflect.Descriptor instead.
func (*HttpBody_Status) Descriptor() ([]byte, []int) {
	return file_leo_status_http_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HttpBody_Status) GetStatus() code.Code {
	if x != nil {
		return x.Status
	}
	return code.Code(0)
}

func (x *HttpBody_Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HttpBody_Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *HttpBody_Status) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *HttpBody_Status) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_leo_status_http_proto protoreflect.FileDescriptor

var file_leo_status_http_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x08, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x48,
	0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0xb0, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x6c, 0x65,
	0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x65, 0x6f, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_leo_status_http_proto_rawDescOnce sync.Once
	file_leo_status_http_proto_rawDescData = file_leo_status_http_proto_rawDesc
)

func file_leo_status_http_proto_rawDescGZIP() []byte {
	file_leo_status_http_proto_rawDescOnce.Do(func() {
		file_leo_status_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_leo_status_http_proto_rawDescData)
	})
	return file_leo_status_http_proto_rawDescData
}

var file_leo_status_http_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_leo_status_http_proto_goTypes = []interface{}{
	(*HttpBody)(nil),        // 0: leo.status.HttpBody
	(*HttpBody_Status)(nil), // 1: leo.status.HttpBody.Status
	(code.Code)(0),          // 2: google.rpc.Code
	(*anypb.Any)(nil),       // 3: google.protobuf.Any
}
var file_leo_status_http_proto_depIdxs = []int32{
	1, // 0: leo.status.HttpBody.error:type_name -> leo.status.HttpBody.Status
	2, // 1: leo.status.HttpBody.Status.status:type_name -> google.rpc.Code
	3, // 2: leo.status.HttpBody.Status.details:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_leo_status_http_proto_init() }
func file_leo_status_http_proto_init() {
	if File_leo_status_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_leo_status_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpBody); i {
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
		file_leo_status_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpBody_Status); i {
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
			RawDescriptor: file_leo_status_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_leo_status_http_proto_goTypes,
		DependencyIndexes: file_leo_status_http_proto_depIdxs,
		MessageInfos:      file_leo_status_http_proto_msgTypes,
	}.Build()
	File_leo_status_http_proto = out.File
	file_leo_status_http_proto_rawDesc = nil
	file_leo_status_http_proto_goTypes = nil
	file_leo_status_http_proto_depIdxs = nil
}
