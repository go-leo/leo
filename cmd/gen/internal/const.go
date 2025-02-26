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

	GrpcTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/grpc")
	HttpTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/http")
	EndpointPackage      = protogen.GoImportPath("github.com/go-kit/kit/endpoint")
	SdPackage            = protogen.GoImportPath("github.com/go-kit/kit/sd")
	LbPackage            = protogen.GoImportPath("github.com/go-kit/kit/sd/lb")
	LogPackage           = protogen.GoImportPath("github.com/go-kit/log")

	EndpointxPackage       = protogen.GoImportPath("github.com/go-leo/leo/v3/endpointx")
	TransportxPackage      = protogen.GoImportPath("github.com/go-leo/leo/v3/transportx")
	GrpcxTransportxPackage = protogen.GoImportPath("github.com/go-leo/leo/v3/transportx/grpcx")
	CqrsPackage            = protogen.GoImportPath("github.com/go-leo/leo/v3/cqrs")
	MetadataxPackage       = protogen.GoImportPath("github.com/go-leo/leo/v3/metadatax")
	StatusxPackage         = protogen.GoImportPath("github.com/go-leo/leo/v3/statusx")
	SdxPackage             = protogen.GoImportPath("github.com/go-leo/leo/v3/sdx")
	LbxPackage             = protogen.GoImportPath("github.com/go-leo/leo/v3/sdx/lbx")
	StainPackage           = protogen.GoImportPath("github.com/go-leo/leo/v3/sdx/stain")

	ErrorxPackage    = protogen.GoImportPath("github.com/go-leo/gox/errorx")
	ConvxPackage     = protogen.GoImportPath("github.com/go-leo/gox/convx")
	StrconvxPackage  = protogen.GoImportPath("github.com/go-leo/gox/strconvx")
	ProtoxPackage    = protogen.GoImportPath("github.com/go-leo/gox/protox")
	UrlxPackage      = protogen.GoImportPath("github.com/go-leo/gox/netx/urlx")
	JsonxPackage     = protogen.GoImportPath("github.com/go-leo/gox/encodingx/jsonx")
	LazyLoadxPackage = protogen.GoImportPath("github.com/go-leo/gox/syncx/lazyloadx")

	GrpcPackage         = protogen.GoImportPath("google.golang.org/grpc")
	GrpcStatusPackage   = protogen.GoImportPath("google.golang.org/grpc/status")
	GrpcMetadataPackage = protogen.GoImportPath("google.golang.org/grpc/metadata")

	RpcHttpPackage = protogen.GoImportPath("google.golang.org/genproto/googleapis/rpc/http")

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

var (
	ProtoJsonPackage               = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	ProtoJsonMarshalOptionsIdent   = ProtoJsonPackage.Ident("MarshalOptions")
	ProtoJsonUnmarshalOptionsIdent = ProtoJsonPackage.Ident("UnmarshalOptions")
)

var (
	HttpxTransportxPackage = protogen.GoImportPath("github.com/go-leo/leo/v3/transportx/httpx")

	OptionIdent     = HttpxTransportxPackage.Ident("Option")
	NewOptionsIdent = HttpxTransportxPackage.Ident("NewOptions")
)
var (
	HttpxCoderPackage = protogen.GoImportPath("github.com/go-leo/leo/v3/transportx/httpx/coder")

	ResponseTransformer = HttpxCoderPackage.Ident("ResponseTransformer")

	DecodeForm                   = HttpxCoderPackage.Ident("DecodeForm")
	DecodeHttpBodyFromRequest    = HttpxCoderPackage.Ident("DecodeHttpBodyFromRequest")
	DecodeHttpRequestFromRequest = HttpxCoderPackage.Ident("DecodeHttpRequestFromRequest")

	EncodeMessageToRequest  = HttpxCoderPackage.Ident("EncodeMessageToRequest")
	EncodeHttpBodyToRequest = HttpxCoderPackage.Ident("EncodeHttpBodyToRequest")

	EncodeMessageToResponse      = HttpxCoderPackage.Ident("EncodeMessageToResponse")
	EncodeHttpBodyToResponse     = HttpxCoderPackage.Ident("EncodeHttpBodyToResponse")
	EncodeHttpResponseToResponse = HttpxCoderPackage.Ident("EncodeHttpResponseToResponse")

	DecodeMessageFromRequest = HttpxCoderPackage.Ident("DecodeMessageFromRequest")

	DecodeMessageFromResponse      = HttpxCoderPackage.Ident("DecodeMessageFromResponse")
	DecodeHttpBodyFromResponse     = HttpxCoderPackage.Ident("DecodeHttpBodyFromResponse")
	DecodeHttpResponseFromResponse = HttpxCoderPackage.Ident("DecodeHttpResponseFromResponse")
)

var (
	MuxPackage = protogen.GoImportPath("github.com/gorilla/mux")
	Router     = MuxPackage.Ident("Router")
	VarsIdent  = MuxPackage.Ident("Vars")
)
