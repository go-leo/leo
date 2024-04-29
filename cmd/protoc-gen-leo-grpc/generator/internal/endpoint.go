package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

type Endpoint struct {
	method *protogen.Method
}

func (e Endpoint) Name() string {
	return e.method.GoName
}

func (e Endpoint) UnexportedName() string {
	s := e.method.GoName
	return strings.ToLower(s[:1]) + s[1:]
}

func (e Endpoint) ArgsName() string {
	return e.method.GoName + "Args"
}

func (e Endpoint) ResName() string {
	return e.method.GoName + "Res"
}

func (e Endpoint) RequestName() string {
	return e.method.GoName + "Request"
}

func (e Endpoint) ResponseName() string {
	return e.method.GoName + "Response"
}

func (e Endpoint) IsStreamingServer() bool {
	return e.method.Desc.IsStreamingServer()
}

func (e Endpoint) InputGoIdent() protogen.GoIdent {
	return e.method.Input.GoIdent
}

func (e Endpoint) OutputGoIdent() protogen.GoIdent {
	return e.method.Output.GoIdent
}

func (e Endpoint) ServerStreamName() string {
	method := e.method
	return method.Parent.GoName + "_" + method.GoName + "Server"
}

func NewEndpoint(method *protogen.Method) *Endpoint {
	return &Endpoint{method: method}
}
