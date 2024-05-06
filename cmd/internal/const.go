package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
	"regexp"
)

var (
	ContextPackage       = protogen.GoImportPath("context")
	UrlPackage           = protogen.GoImportPath("net/url")
	GrpcTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/grpc")
	HttpTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/http")
	EndpointPackage      = protogen.GoImportPath("github.com/go-kit/kit/endpoint")
	GrpcPackage          = protogen.GoImportPath("google.golang.org/grpc")
	EndpointxPackage     = protogen.GoImportPath("github.com/go-leo/leo/v3/endpointx")
	StatusPackage        = protogen.GoImportPath("google.golang.org/grpc/status")
	MuxPackage           = protogen.GoImportPath("github.com/gorilla/mux")
	HttpPackage          = protogen.GoImportPath("net/http")
	FmtPackage           = protogen.GoImportPath("fmt")
	StrconvPackage       = protogen.GoImportPath("strconv")
	ConvxPackage         = protogen.GoImportPath("github.com/go-leo/gox/convx")
	IOPackage            = protogen.GoImportPath("io")
	ProtoJsonPackage     = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	RpcHttpPackage       = protogen.GoImportPath("google.golang.org/genproto/googleapis/rpc/http")
	ProtoPackage         = protogen.GoImportPath("google.golang.org/protobuf/proto")
	WrapperspbPackage    = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	FieldmaskpbPackage   = protogen.GoImportPath("google.golang.org/protobuf/types/known/fieldmaskpb")
	BytesPackage         = protogen.GoImportPath("bytes")
	StringsPackage       = protogen.GoImportPath("strings")
	JsonPackage          = protogen.GoImportPath("encoding/json")
)

var (
	namedPathPattern = regexp.MustCompile("{(.+)=(.+)}")
	pathPattern      = regexp.MustCompile("{([^=}]+)}")
)
