package internal

import (
	"fmt"
	"github.com/go-leo/gox/slicex"
	"github.com/go-leo/leo/v3/cqrs"
	"golang.org/x/exp/slices"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"net/http"
	"strings"
)

type Endpoint struct {
	method         *protogen.Method
	httpRule       *HttpRule
	responsibility cqrs.Responsibility
}

func (e Endpoint) Name() string {
	return e.method.GoName
}

func (e Endpoint) UnexportedName() string {
	s := e.Name()
	return strings.ToLower(s[:1]) + s[1:]
}

func (e Endpoint) FullName() string {
	return fmt.Sprintf("/%s/%s", e.method.Parent.Desc.FullName(), e.method.Desc.Name())
}

func (e Endpoint) RequestDecoderName() string {
	return fmt.Sprintf("_%s_%s_RequestDecoder", e.method.Parent.GoName, e.method.GoName)
}

func (e Endpoint) ResponseEncoderName() string {
	return fmt.Sprintf("_%s_%s_ResponseEncoder", e.method.Parent.GoName, e.method.GoName)
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

func (e Endpoint) IsStreaming() bool {
	return e.method.Desc.IsStreamingServer() || e.method.Desc.IsStreamingClient()
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

func (e Endpoint) Responsibility() cqrs.Responsibility {
	return e.responsibility
}

func (e Endpoint) IsCommand() bool {
	return e.responsibility == cqrs.Responsibility_Command
}

func (e Endpoint) IsQuery() bool {
	return e.responsibility == cqrs.Responsibility_Query
}

func (e Endpoint) HttpRule() *HttpRule {
	return e.httpRule
}

func (e Endpoint) ParseParameters() (*protogen.Message, *protogen.Field, []*protogen.Field, []*protogen.Field, []*protogen.Field, error) {
	httpRule := e.httpRule
	bodyParameter := httpRule.Body()
	path, namedPathName, _, namedPathParameters := httpRule.RegularizePath(httpRule.Path())
	pathParameters := slicex.Difference(httpRule.PathParameters(path), namedPathParameters)

	// body arguments
	var bodyMessage *protogen.Message
	var bodyField *protogen.Field
	switch bodyParameter {
	case "":
		// ignore
	case "*":
		bodyMessage = e.Input()
	default:
		bodyField = FindField(bodyParameter, e.Input())
		if bodyField == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("%s, failed to find body field %s", e.FullName(), bodyParameter)
		}
	}

	// namedPathParameters
	var namedPathFields []*protogen.Field
	input := e.Input()
	if len(namedPathName) > 0 {
		namedPathParameters := strings.Split(namedPathName, ".")
		for i, namedPathParameter := range namedPathParameters {
			field := FindField(namedPathParameter, input)
			if field == nil {
				return nil, nil, nil, nil, nil, fmt.Errorf("%s, failed to find named path field %s", e.FullName(), namedPathName)
			}
			if i < len(namedPathParameters)-1 {
				if field.Desc.Kind() != protoreflect.MessageKind {
					return nil, nil, nil, nil, nil, fmt.Errorf("%s, %s is not message", e.FullName(), field.Desc.Name())
				}
			} else {
				switch field.Desc.Kind() {
				case protoreflect.StringKind:
				case protoreflect.MessageKind:
					switch field.Message.Desc.FullName() {
					case "google.protobuf.StringValue":
					default:
						return nil, nil, nil, nil, nil, fmt.Errorf("%s, named path parameters do not support %s", field.Message.Desc.FullName())
					}
				default:
					return nil, nil, nil, nil, nil, fmt.Errorf("%s, named path parameters do not support %s", field.Desc.Kind())
				}
			}
			namedPathFields = append(namedPathFields, field)
			input = field.Message
		}
	}

	var pathFields []*protogen.Field
	for _, pathParameter := range pathParameters {
		field := FindField(pathParameter, e.Input())
		if field == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("%s, failed to find path field %s", e.FullName(), bodyParameter)
		}
		if field.Desc.IsList() || field.Desc.IsMap() {
			return nil, nil, nil, nil, nil, fmt.Errorf("%s, path parameters do not support list or map", e.FullName())
		}

		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
		case protoreflect.FloatKind: // float32
		case protoreflect.DoubleKind: // float64
		case protoreflect.StringKind: // string
		case protoreflect.EnumKind: // enum
		case protoreflect.MessageKind:
			message := field.Message
			switch message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
			case "google.protobuf.FloatValue":
			case "google.protobuf.Int64Value":
			case "google.protobuf.UInt64Value":
			case "google.protobuf.Int32Value":
			case "google.protobuf.UInt32Value":
			case "google.protobuf.BoolValue":
			case "google.protobuf.StringValue":
			default:
				return nil, nil, nil, nil, nil, fmt.Errorf("%s, path parameters do not support %s", e.FullName(), message.Desc.FullName())
			}
		default:
			return nil, nil, nil, nil, nil, fmt.Errorf("%s, path parameters do not support %s", e.FullName(), field.Desc.Kind())
		}

		pathFields = append(pathFields, field)
	}

	var queryFields []*protogen.Field
	if bodyMessage != nil {
		return bodyMessage, bodyField, namedPathFields, pathFields, queryFields, nil
	}
	for _, field := range e.Input().Fields {
		if field == bodyField {
			continue
		}
		if slices.Contains(namedPathFields, field) {
			continue
		}
		if slices.Contains(namedPathFields, field) {
			continue
		}
		if slices.Contains(pathFields, field) {
			continue
		}
		if field.Desc.IsMap() {
			continue
		}
		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
		case protoreflect.FloatKind: // float32
		case protoreflect.DoubleKind: // float64
		case protoreflect.StringKind: // string
		case protoreflect.EnumKind: // enum
		case protoreflect.MessageKind:
			message := field.Message
			switch message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
			case "google.protobuf.FloatValue":
			case "google.protobuf.Int64Value":
			case "google.protobuf.UInt64Value":
			case "google.protobuf.Int32Value":
			case "google.protobuf.UInt32Value":
			case "google.protobuf.BoolValue":
			case "google.protobuf.StringValue":
			default:
				continue
			}
		default:
			continue
		}
		queryFields = append(queryFields, field)
	}
	return bodyMessage, bodyField, namedPathFields, pathFields, queryFields, nil
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

func (r *HttpRule) RegularizePath(path string) (string, string, string, []string) {
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
	return path, name, template, parameters
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
