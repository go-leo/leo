package internal

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"net/http"
	"strings"
)

type Endpoint struct {
	method    *protogen.Method
	httpRules []*HttpRule
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

func (e Endpoint) HttpRules() []*HttpRule {
	return e.httpRules
}

func NewEndpoint(method *protogen.Method, rawRules []*annotations.HttpRule) *Endpoint {
	rules := make([]*HttpRule, 0, len(rawRules))
	for _, rule := range rawRules {
		rules = append(rules, &HttpRule{rule: rule})
	}
	return &Endpoint{method: method, httpRules: rules}
}

type HttpRule struct {
	rule *annotations.HttpRule
}

func (r *HttpRule) Method() string {
	switch pattern := r.rule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		return http.MethodGet
	case *annotations.HttpRule_Post:
		return http.MethodPost
	case *annotations.HttpRule_Put:
		return http.MethodPut
	case *annotations.HttpRule_Delete:
		return http.MethodDelete
	case *annotations.HttpRule_Patch:
		return http.MethodPatch
	case *annotations.HttpRule_Custom:
		return pattern.Custom.GetKind()
	default:
		return ""
	}
}

func (r *HttpRule) Path() string {
	switch pattern := r.rule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		return pattern.Get
	case *annotations.HttpRule_Post:
		return pattern.Post
	case *annotations.HttpRule_Put:
		return pattern.Put
	case *annotations.HttpRule_Delete:
		return pattern.Delete
	case *annotations.HttpRule_Patch:
		return pattern.Patch
	case *annotations.HttpRule_Custom:
		return pattern.Custom.GetPath()
	default:
		return ""
	}
}
