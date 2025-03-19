package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
	"regexp"
)

var (
	ContextPackage = protogen.GoImportPath("context")
	Context        = ContextPackage.Ident("Context")
	TODO           = ContextPackage.Ident("TODO")
)

var (
	UrlPackage     = protogen.GoImportPath("net/url")
	ErrorsPackage  = protogen.GoImportPath("errors")
	TimePackage    = protogen.GoImportPath("time")
	StrconvPackage = protogen.GoImportPath("strconv")
	IOPackage      = protogen.GoImportPath("io")
	StringsPackage = protogen.GoImportPath("strings")
	JsonPackage    = protogen.GoImportPath("encoding/json")

	GrpcTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/grpc")
	HttpTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/http")
	EndpointPackage      = protogen.GoImportPath("github.com/go-kit/kit/endpoint")
	SdPackage            = protogen.GoImportPath("github.com/go-kit/kit/sd")
	LbPackage            = protogen.GoImportPath("github.com/go-kit/kit/sd/lb")
	LogPackage           = protogen.GoImportPath("github.com/go-kit/log")

	TransportxPackage      = protogen.GoImportPath("github.com/go-leo/leo/v3/transportx")
	GrpcxTransportxPackage = protogen.GoImportPath("github.com/go-leo/leo/v3/transportx/grpctransportx")
	StatusxPackage         = protogen.GoImportPath("github.com/go-leo/leo/v3/statusx")
	SdxPackage             = protogen.GoImportPath("github.com/go-leo/leo/v3/sdx")
	LbxPackage             = protogen.GoImportPath("github.com/go-leo/leo/v3/sdx/lbx")

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
	BytesPackage = protogen.GoImportPath("bytes")
	Buffer       = BytesPackage.Ident("Buffer")
)

var (
	EndpointxPackage        = protogen.GoImportPath("github.com/go-leo/leo/v3/endpointx")
	EndpointxChain          = EndpointxPackage.Ident("Chain")
	EndpointxNameInjector   = EndpointxPackage.Ident("NameInjector")
	EndpointxHandleEndpoint = EndpointxPackage.Ident("HandleEndpoint")
)

var (
	MetadataxPackage              = protogen.GoImportPath("github.com/go-leo/leo/v3/metadatax")
	MetadataxMetadata             = MetadataxPackage.Ident("Metadata")
	MetadataxGrpcOutgoingInjector = MetadataxPackage.Ident("GrpcOutgoingInjector")
	MetadataxGrpcIncomingInjector = MetadataxPackage.Ident("GrpcIncomingInjector")
	MetadataxHttpOutgoingInjector = MetadataxPackage.Ident("HttpOutgoingInjector")
	MetadataxHttpIncomingInjector = MetadataxPackage.Ident("HttpIncomingInjector")
)

var (
	HttpPackage           = protogen.GoImportPath("net/http")
	Handler               = HttpPackage.Ident("Handler")
	Request               = HttpPackage.Ident("Request")
	Response              = HttpPackage.Ident("Response")
	ResponseWriter        = HttpPackage.Ident("ResponseWriter")
	Header                = HttpPackage.Ident("Header")
	NewRequestWithContext = HttpPackage.Ident("NewRequestWithContext")
	StatusOK              = HttpPackage.Ident("StatusOK")
	MethodGet             = HttpPackage.Ident("MethodGet")
	MethodPost            = HttpPackage.Ident("MethodPost")
	MethodPut             = HttpPackage.Ident("MethodPut")
	MethodDelete          = HttpPackage.Ident("MethodDelete")
	MethodPatch           = HttpPackage.Ident("MethodPatch")
)

var (
	FmtPackage = protogen.GoImportPath("fmt")
	Errorf     = FmtPackage.Ident("Errorf")
	Sprintf    = FmtPackage.Ident("Sprintf")
)

var (
	samplePathPattern = regexp.MustCompile("{([^=}]+)}")
	namedPathPattern  = regexp.MustCompile("{([^{}]+)=([^{}]+)}")
)

var (
	ProtoJsonPackage               = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	ProtoJsonMarshalOptionsIdent   = ProtoJsonPackage.Ident("MarshalOptions")
	ProtoJsonUnmarshalOptionsIdent = ProtoJsonPackage.Ident("UnmarshalOptions")
)

var (
	StainxPackage              = protogen.GoImportPath("github.com/go-leo/leo/v3/stainx")
	StainxGrpcOutgoingInjector = StainxPackage.Ident("GrpcOutgoingInjector")
	StainxGrpcIncomingInjector = StainxPackage.Ident("GrpcIncomingInjector")
	StainxHttpOutgoingInjector = StainxPackage.Ident("HttpOutgoingInjector")
	StainxHttpIncomingInjector = StainxPackage.Ident("HttpIncomingInjector")
	StainxColorExtractor       = StainxPackage.Ident("ColorExtractor")
)

var (
	HttpxTransportxPackage = protogen.GoImportPath("github.com/go-leo/leo/v3/transportx/httptransportx")

	OptionIdent      = HttpxTransportxPackage.Ident("Option")
	NewOptionsIdent  = HttpxTransportxPackage.Ident("NewOptions")
	ServerOption     = HttpxTransportxPackage.Ident("ServerOption")
	ServerOptions    = HttpxTransportxPackage.Ident("ServerOptions")
	NewServerOptions = HttpxTransportxPackage.Ident("NewServerOptions")
)

var (
	TimeoutxPackage = protogen.GoImportPath("github.com/go-leo/leo/v3/timeoutx")

	TimeoutxOutgoingInjector = TimeoutxPackage.Ident("OutgoingInjector")
	TimeoutxIncomingInjector = TimeoutxPackage.Ident("IncomingInjector")
	TimeoutxCancelInvoker    = TimeoutxPackage.Ident("CancelInvoker")
)

var (
	HttpxCoderPackage = protogen.GoImportPath("github.com/go-leo/leo/v3/transportx/httptransportx/coder")

	DecodeForm                   = HttpxCoderPackage.Ident("DecodeForm")
	DecodeMessageFromRequest     = HttpxCoderPackage.Ident("DecodeMessageFromRequest")
	DecodeHttpBodyFromRequest    = HttpxCoderPackage.Ident("DecodeHttpBodyFromRequest")
	DecodeHttpRequestFromRequest = HttpxCoderPackage.Ident("DecodeHttpRequestFromRequest")

	EncodeMessageToRequest     = HttpxCoderPackage.Ident("EncodeMessageToRequest")
	EncodeHttpBodyToRequest    = HttpxCoderPackage.Ident("EncodeHttpBodyToRequest")
	EncodeHttpRequestToRequest = HttpxCoderPackage.Ident("EncodeHttpRequestToRequest")

	DecodeErrorFromResponse        = HttpxCoderPackage.Ident("DecodeErrorFromResponse")
	DecodeMessageFromResponse      = HttpxCoderPackage.Ident("DecodeMessageFromResponse")
	DecodeHttpBodyFromResponse     = HttpxCoderPackage.Ident("DecodeHttpBodyFromResponse")
	DecodeHttpResponseFromResponse = HttpxCoderPackage.Ident("DecodeHttpResponseFromResponse")

	EncodeMessageToResponse      = HttpxCoderPackage.Ident("EncodeMessageToResponse")
	EncodeHttpBodyToResponse     = HttpxCoderPackage.Ident("EncodeHttpBodyToResponse")
	EncodeHttpResponseToResponse = HttpxCoderPackage.Ident("EncodeHttpResponseToResponse")
	EncodeErrorToResponse        = HttpxCoderPackage.Ident("EncodeErrorToResponse")
)

var (
	MuxPackage = protogen.GoImportPath("github.com/gorilla/mux")
	Router     = MuxPackage.Ident("Router")
	VarsIdent  = MuxPackage.Ident("Vars")
)

var (
	HttpxPackage = protogen.GoImportPath("github.com/go-leo/gox/netx/httpx")
	CopyHeader   = HttpxPackage.Ident("CopyHeader")
)

var (
	CqrsPackage    = protogen.GoImportPath("github.com/go-leo/leo/v3/cqrs")
	QueryHandler   = CqrsPackage.Ident("QueryHandler")
	CommandHandler = CqrsPackage.Ident("CommandHandler")
	Bus            = CqrsPackage.Ident("Bus")
)
