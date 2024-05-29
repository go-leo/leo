package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
	"regexp"
)

var (
	ContextPackage = protogen.GoImportPath("context")
	UrlPackage     = protogen.GoImportPath("net/url")
	HttpPackage    = protogen.GoImportPath("net/http")
	FmtPackage     = protogen.GoImportPath("fmt")
	ErrorsPackage  = protogen.GoImportPath("errors")
	TimePackage    = protogen.GoImportPath("time")
	StrconvPackage = protogen.GoImportPath("strconv")
	IOPackage      = protogen.GoImportPath("io")
	BytesPackage   = protogen.GoImportPath("bytes")
	StringsPackage = protogen.GoImportPath("strings")
	JsonPackage    = protogen.GoImportPath("encoding/json")

	MuxPackage = protogen.GoImportPath("github.com/gorilla/mux")

	GrpcTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/grpc")
	HttpTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/http")
	EndpointPackage      = protogen.GoImportPath("github.com/go-kit/kit/endpoint")

	EndpointxPackage = protogen.GoImportPath("github.com/go-leo/leo/v3/endpointx")
	CqrsPackage      = protogen.GoImportPath("github.com/go-leo/leo/v3/cqrs")
	MetadataxPackage = protogen.GoImportPath("github.com/go-leo/leo/v3/metadatax")

	ErrorxPackage   = protogen.GoImportPath("github.com/go-leo/gox/errorx")
	ConvxPackage    = protogen.GoImportPath("github.com/go-leo/gox/convx")
	StrconvxPackage = protogen.GoImportPath("github.com/go-leo/gox/strconvx")
	ProtoxPackage   = protogen.GoImportPath("github.com/go-leo/gox/protox")
	UrlxPackage     = protogen.GoImportPath("github.com/go-leo/gox/netx/urlx")
	JsonxPackage    = protogen.GoImportPath("github.com/go-leo/gox/encodingx/jsonx")

	GrpcPackage   = protogen.GoImportPath("google.golang.org/grpc")
	StatusPackage = protogen.GoImportPath("google.golang.org/grpc/status")

	RpcHttpPackage = protogen.GoImportPath("google.golang.org/genproto/googleapis/rpc/http")

	ProtoJsonPackage   = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	ProtoPackage       = protogen.GoImportPath("google.golang.org/protobuf/proto")
	WrapperspbPackage  = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	FieldmaskpbPackage = protogen.GoImportPath("google.golang.org/protobuf/types/known/fieldmaskpb")
	StructpbPackage    = protogen.GoImportPath("google.golang.org/protobuf/types/known/structpb")
	AnypbPackage       = protogen.GoImportPath("google.golang.org/protobuf/types/known/anypb")
)

var (
	namedPathPattern = regexp.MustCompile("{([^{}]+)=([^{}]+)}")
	pathPattern      = regexp.MustCompile("{([^=}]+)}")
)

var (
	ContentTypeKey  = "Content-Type"
	JsonContentType = "application/json; charset=utf-8"
)
