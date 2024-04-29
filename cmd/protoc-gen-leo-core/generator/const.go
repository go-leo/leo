package generator

import "google.golang.org/protobuf/compiler/protogen"

var (
	contextPackage       = protogen.GoImportPath("context")
	grpcTransportPackage = protogen.GoImportPath("github.com/go-kit/kit/transport/grpc")
	endpointPackage      = protogen.GoImportPath("github.com/go-kit/kit/endpoint")
	grpcPackage          = protogen.GoImportPath("google.golang.org/grpc")
	endpointxPackage     = protogen.GoImportPath("github.com/go-leo/kitx/endpointx")
	statusPackage        = protogen.GoImportPath("google.golang.org/grpc/status")
)
