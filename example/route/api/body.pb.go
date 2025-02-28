// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: api/body.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	http "google.golang.org/genproto/googleapis/rpc/http"
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

type BodyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *BodyRequest_User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *BodyRequest) Reset() {
	*x = BodyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_body_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyRequest) ProtoMessage() {}

func (x *BodyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_body_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyRequest.ProtoReflect.Descriptor instead.
func (*BodyRequest) Descriptor() ([]byte, []int) {
	return file_api_body_proto_rawDescGZIP(), []int{0}
}

func (x *BodyRequest) GetUser() *BodyRequest_User {
	if x != nil {
		return x.User
	}
	return nil
}

type HttpBodyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *httpbody.HttpBody `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *HttpBodyRequest) Reset() {
	*x = HttpBodyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_body_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpBodyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpBodyRequest) ProtoMessage() {}

func (x *HttpBodyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_body_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpBodyRequest.ProtoReflect.Descriptor instead.
func (*HttpBodyRequest) Descriptor() ([]byte, []int) {
	return file_api_body_proto_rawDescGZIP(), []int{1}
}

func (x *HttpBodyRequest) GetBody() *httpbody.HttpBody {
	if x != nil {
		return x.Body
	}
	return nil
}

type BodyRequest_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone   string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *BodyRequest_User) Reset() {
	*x = BodyRequest_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_body_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyRequest_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyRequest_User) ProtoMessage() {}

func (x *BodyRequest_User) ProtoReflect() protoreflect.Message {
	mi := &file_api_body_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyRequest_User.ProtoReflect.Descriptor instead.
func (*BodyRequest_User) Descriptor() ([]byte, []int) {
	return file_api_body_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BodyRequest_User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BodyRequest_User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *BodyRequest_User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *BodyRequest_User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_api_body_proto protoreflect.FileDescriptor

var file_api_body_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x60, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3b, 0x0a, 0x0f, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x32, 0xe4, 0x04, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x61, 0x0a, 0x08, 0x53,
	0x74, 0x61, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x23, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x6f, 0x64, 0x79,
	0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x66,
	0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x23, 0x2e, 0x6c, 0x65,
	0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x62, 0x6f, 0x64, 0x79, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65,
	0x64, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x50, 0x0a, 0x07, 0x4e, 0x6f, 0x6e, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x64, 0x0a, 0x10, 0x48, 0x74, 0x74, 0x70,
	0x42, 0x6f, 0x64, 0x79, 0x53, 0x74, 0x61, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f,
	0x64, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x1a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x62, 0x6f, 0x64, 0x79, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x7c,
	0x0a, 0x11, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x27, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x32, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x62, 0x6f, 0x64,
	0x79, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x5b, 0x0a, 0x0b,
	0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x6c,
	0x65, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_body_proto_rawDescOnce sync.Once
	file_api_body_proto_rawDescData = file_api_body_proto_rawDesc
)

func file_api_body_proto_rawDescGZIP() []byte {
	file_api_body_proto_rawDescOnce.Do(func() {
		file_api_body_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_body_proto_rawDescData)
	})
	return file_api_body_proto_rawDescData
}

var file_api_body_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_body_proto_goTypes = []interface{}{
	(*BodyRequest)(nil),       // 0: leo.example.route.body.BodyRequest
	(*HttpBodyRequest)(nil),   // 1: leo.example.route.body.HttpBodyRequest
	(*BodyRequest_User)(nil),  // 2: leo.example.route.body.BodyRequest.User
	(*httpbody.HttpBody)(nil), // 3: google.api.HttpBody
	(*emptypb.Empty)(nil),     // 4: google.protobuf.Empty
	(*http.HttpRequest)(nil),  // 5: google.rpc.HttpRequest
}
var file_api_body_proto_depIdxs = []int32{
	2, // 0: leo.example.route.body.BodyRequest.user:type_name -> leo.example.route.body.BodyRequest.User
	3, // 1: leo.example.route.body.HttpBodyRequest.body:type_name -> google.api.HttpBody
	0, // 2: leo.example.route.body.Body.StarBody:input_type -> leo.example.route.body.BodyRequest
	0, // 3: leo.example.route.body.Body.NamedBody:input_type -> leo.example.route.body.BodyRequest
	4, // 4: leo.example.route.body.Body.NonBody:input_type -> google.protobuf.Empty
	3, // 5: leo.example.route.body.Body.HttpBodyStarBody:input_type -> google.api.HttpBody
	1, // 6: leo.example.route.body.Body.HttpBodyNamedBody:input_type -> leo.example.route.body.HttpBodyRequest
	5, // 7: leo.example.route.body.Body.HttpRequest:input_type -> google.rpc.HttpRequest
	4, // 8: leo.example.route.body.Body.StarBody:output_type -> google.protobuf.Empty
	4, // 9: leo.example.route.body.Body.NamedBody:output_type -> google.protobuf.Empty
	4, // 10: leo.example.route.body.Body.NonBody:output_type -> google.protobuf.Empty
	4, // 11: leo.example.route.body.Body.HttpBodyStarBody:output_type -> google.protobuf.Empty
	4, // 12: leo.example.route.body.Body.HttpBodyNamedBody:output_type -> google.protobuf.Empty
	4, // 13: leo.example.route.body.Body.HttpRequest:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_body_proto_init() }
func file_api_body_proto_init() {
	if File_api_body_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_body_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyRequest); i {
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
		file_api_body_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpBodyRequest); i {
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
		file_api_body_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyRequest_User); i {
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
			RawDescriptor: file_api_body_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_body_proto_goTypes,
		DependencyIndexes: file_api_body_proto_depIdxs,
		MessageInfos:      file_api_body_proto_msgTypes,
	}.Build()
	File_api_body_proto = out.File
	file_api_body_proto_rawDesc = nil
	file_api_body_proto_goTypes = nil
	file_api_body_proto_depIdxs = nil
}
