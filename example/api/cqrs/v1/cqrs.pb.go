// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: cqrs/v1/cqrs.proto

package cqrs

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cqrs_v1_cqrs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cqrs_v1_cqrs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_cqrs_v1_cqrs_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *QueryReply) Reset() {
	*x = QueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cqrs_v1_cqrs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReply) ProtoMessage() {}

func (x *QueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_cqrs_v1_cqrs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReply.ProtoReflect.Descriptor instead.
func (*QueryReply) Descriptor() ([]byte, []int) {
	return file_cqrs_v1_cqrs_proto_rawDescGZIP(), []int{1}
}

func (x *QueryReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cqrs_v1_cqrs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cqrs_v1_cqrs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_cqrs_v1_cqrs_proto_rawDescGZIP(), []int{2}
}

func (x *CommandRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CommandReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommandReply) Reset() {
	*x = CommandReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cqrs_v1_cqrs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandReply) ProtoMessage() {}

func (x *CommandReply) ProtoReflect() protoreflect.Message {
	mi := &file_cqrs_v1_cqrs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandReply.ProtoReflect.Descriptor instead.
func (*CommandReply) Descriptor() ([]byte, []int) {
	return file_cqrs_v1_cqrs_proto_rawDescGZIP(), []int{3}
}

type QueryOneOfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueryOneOfRequest) Reset() {
	*x = QueryOneOfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cqrs_v1_cqrs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOneOfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOneOfRequest) ProtoMessage() {}

func (x *QueryOneOfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cqrs_v1_cqrs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOneOfRequest.ProtoReflect.Descriptor instead.
func (*QueryOneOfRequest) Descriptor() ([]byte, []int) {
	return file_cqrs_v1_cqrs_proto_rawDescGZIP(), []int{4}
}

func (x *QueryOneOfRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QueryOneOfReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*QueryOneOfReply_Name
	//	*QueryOneOfReply_Id
	Data isQueryOneOfReply_Data `protobuf_oneof:"data"`
}

func (x *QueryOneOfReply) Reset() {
	*x = QueryOneOfReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cqrs_v1_cqrs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOneOfReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOneOfReply) ProtoMessage() {}

func (x *QueryOneOfReply) ProtoReflect() protoreflect.Message {
	mi := &file_cqrs_v1_cqrs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOneOfReply.ProtoReflect.Descriptor instead.
func (*QueryOneOfReply) Descriptor() ([]byte, []int) {
	return file_cqrs_v1_cqrs_proto_rawDescGZIP(), []int{5}
}

func (m *QueryOneOfReply) GetData() isQueryOneOfReply_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *QueryOneOfReply) GetName() string {
	if x, ok := x.GetData().(*QueryOneOfReply_Name); ok {
		return x.Name
	}
	return ""
}

func (x *QueryOneOfReply) GetId() string {
	if x, ok := x.GetData().(*QueryOneOfReply_Id); ok {
		return x.Id
	}
	return ""
}

type isQueryOneOfReply_Data interface {
	isQueryOneOfReply_Data()
}

type QueryOneOfReply_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type QueryOneOfReply_Id struct {
	Id string `protobuf:"bytes,2,opt,name=id,proto3,oneof"`
}

func (*QueryOneOfReply_Name) isQueryOneOfReply_Data() {}

func (*QueryOneOfReply_Id) isQueryOneOfReply_Data() {}

var File_cqrs_v1_cqrs_proto protoreflect.FileDescriptor

var file_cqrs_v1_cqrs_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x71, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x27, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x6e, 0x65, 0x4f,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x0f,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0xab, 0x03, 0x0a, 0x04, 0x43, 0x71, 0x72, 0x73, 0x12, 0x60, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x1e, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x63, 0x71, 0x72, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x63, 0x71, 0x72, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x71, 0x72, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x68, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a,
	0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x70, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x6e, 0x65,
	0x4f, 0x66, 0x12, 0x1e, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x6e, 0x65, 0x4f, 0x66,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a,
	0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2d, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x65, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x20, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x71, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6c,
	0x65, 0x6f, 0x2f, 0x6c, 0x65, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x71,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cqrs_v1_cqrs_proto_rawDescOnce sync.Once
	file_cqrs_v1_cqrs_proto_rawDescData = file_cqrs_v1_cqrs_proto_rawDesc
)

func file_cqrs_v1_cqrs_proto_rawDescGZIP() []byte {
	file_cqrs_v1_cqrs_proto_rawDescOnce.Do(func() {
		file_cqrs_v1_cqrs_proto_rawDescData = protoimpl.X.CompressGZIP(file_cqrs_v1_cqrs_proto_rawDescData)
	})
	return file_cqrs_v1_cqrs_proto_rawDescData
}

var file_cqrs_v1_cqrs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cqrs_v1_cqrs_proto_goTypes = []interface{}{
	(*QueryRequest)(nil),      // 0: leo.example.cqrs.QueryRequest
	(*QueryReply)(nil),        // 1: leo.example.cqrs.QueryReply
	(*CommandRequest)(nil),    // 2: leo.example.cqrs.CommandRequest
	(*CommandReply)(nil),      // 3: leo.example.cqrs.CommandReply
	(*QueryOneOfRequest)(nil), // 4: leo.example.cqrs.QueryOneOfRequest
	(*QueryOneOfReply)(nil),   // 5: leo.example.cqrs.QueryOneOfReply
	(*emptypb.Empty)(nil),     // 6: google.protobuf.Empty
}
var file_cqrs_v1_cqrs_proto_depIdxs = []int32{
	0, // 0: leo.example.cqrs.Cqrs.Query:input_type -> leo.example.cqrs.QueryRequest
	2, // 1: leo.example.cqrs.Cqrs.Command:input_type -> leo.example.cqrs.CommandRequest
	0, // 2: leo.example.cqrs.Cqrs.QueryOneOf:input_type -> leo.example.cqrs.QueryRequest
	2, // 3: leo.example.cqrs.Cqrs.CommandEmpty:input_type -> leo.example.cqrs.CommandRequest
	1, // 4: leo.example.cqrs.Cqrs.Query:output_type -> leo.example.cqrs.QueryReply
	3, // 5: leo.example.cqrs.Cqrs.Command:output_type -> leo.example.cqrs.CommandReply
	5, // 6: leo.example.cqrs.Cqrs.QueryOneOf:output_type -> leo.example.cqrs.QueryOneOfReply
	6, // 7: leo.example.cqrs.Cqrs.CommandEmpty:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cqrs_v1_cqrs_proto_init() }
func file_cqrs_v1_cqrs_proto_init() {
	if File_cqrs_v1_cqrs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cqrs_v1_cqrs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_cqrs_v1_cqrs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReply); i {
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
		file_cqrs_v1_cqrs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandRequest); i {
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
		file_cqrs_v1_cqrs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandReply); i {
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
		file_cqrs_v1_cqrs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOneOfRequest); i {
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
		file_cqrs_v1_cqrs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOneOfReply); i {
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
	file_cqrs_v1_cqrs_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*QueryOneOfReply_Name)(nil),
		(*QueryOneOfReply_Id)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cqrs_v1_cqrs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cqrs_v1_cqrs_proto_goTypes,
		DependencyIndexes: file_cqrs_v1_cqrs_proto_depIdxs,
		MessageInfos:      file_cqrs_v1_cqrs_proto_msgTypes,
	}.Build()
	File_cqrs_v1_cqrs_proto = out.File
	file_cqrs_v1_cqrs_proto_rawDesc = nil
	file_cqrs_v1_cqrs_proto_goTypes = nil
	file_cqrs_v1_cqrs_proto_depIdxs = nil
}
