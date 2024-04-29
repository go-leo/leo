package internal

import "google.golang.org/protobuf/compiler/protogen"

var (
	ContextPackage       = protogen.GoImportPath("context")
	GrpcTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/grpc")
	HttpTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/http")
	EndpointPackage      = protogen.GoImportPath("github.com/go-kit/kit/endpoint")
	GrpcPackage          = protogen.GoImportPath("google.golang.org/grpc")
	EndpointxPackage     = protogen.GoImportPath("github.com/go-leo/kitx/endpointx")
	StatusPackage        = protogen.GoImportPath("google.golang.org/grpc/status")
)
