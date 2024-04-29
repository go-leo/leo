package internal

import (
	"fmt"
	"golang.org/x/exp/slices"
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

func (e Endpoint) FullName() string {
	return fmt.Sprintf("/%s/%s", e.method.Parent.Desc.FullName(), e.method.Desc.Name())
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

func (e Endpoint) Input() *protogen.Message {
	return e.method.Input
}

func (e Endpoint) Output() *protogen.Message {
	return e.method.Output
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

func (r *HttpRule) RegularizePath(path string) (string, []string, string, []string) {
	var name string
	var parameters []string
	var template string
	// Find named path parameters like {name=shelves/*}
	if matches := namedPathPattern.FindStringSubmatch(path); matches != nil {
		name = matches[1]
		starredPath := matches[2]
		parts := strings.Split(starredPath, "/")
		newParts := slices.Clone(parts)
		templateParts := slices.Clone(parts)
		// "things/*/otherthings/*" => "things/{thingsId}/otherthings/{otherthingsId}"
		for i := 0; i < len(parts)-1; i += 2 {
			namedPathParameter := singular(newParts[i])
			newParts[i+1] = "{" + namedPathParameter + "}"
			templateParts[i+1] = "%s"
			parameters = append(parameters, namedPathParameter)
		}
		newPath := strings.Join(newParts, "/")
		template = strings.Join(templateParts, "/")
		path = strings.Replace(path, matches[0], newPath, 1)
	}
	return path, strings.Split(name, "."), template, parameters
}

func (r *HttpRule) PathParameters(path string) []string {
	// Find simple path parameters like {id}
	var parameters []string
	if allMatches := pathPattern.FindAllStringSubmatch(path, -1); allMatches != nil {
		for _, matches := range allMatches {
			pathParameter := matches[1]
			parameters = append(parameters, pathParameter)
			path = strings.Replace(path, matches[1], pathParameter, 1)
		}
	}
	return parameters
}

func (r *HttpRule) Body() string {
	return r.rule.GetBody()
}

func (r *HttpRule) ResponseBody() string {
	return r.rule.GetResponseBody()
}
