package internal

import (
	"fmt"
	"github.com/go-leo/gox/slicex"
	"golang.org/x/exp/slices"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

type Endpoint struct {
	service     *Service
	protoMethod *protogen.Method
	httpRule    *annotations.HttpRule

	// http rule pattern
	httpMethod                protogen.GoIdent
	rawPath                   string
	routePath                 string
	namedPathFieldName        string
	namedPathTemplate         string
	namedPathFieldsParameters []string
	pathFieldsParameters      []string

	// request parameters
	bodyMessage     *protogen.Message
	bodyField       *protogen.Field
	namedPathFields []*protogen.Field
	pathFields      []*protogen.Field
	queryFields     []*protogen.Field
}

func (e *Endpoint) Name() string {
	return e.protoMethod.GoName
}

func (e *Endpoint) Unexported(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func (e *Endpoint) FullName() string {
	return fmt.Sprintf("/%s/%s", e.protoMethod.Parent.Desc.FullName(), e.protoMethod.Desc.Name())
}

func (e *Endpoint) IsStreaming() bool {
	return e.protoMethod.Desc.IsStreamingServer() || e.protoMethod.Desc.IsStreamingClient()
}

func (e *Endpoint) Input() *protogen.Message {
	return e.protoMethod.Input
}

func (e *Endpoint) Output() *protogen.Message {
	return e.protoMethod.Output
}

func (e *Endpoint) InputGoIdent() protogen.GoIdent {
	return e.protoMethod.Input.GoIdent
}

func (e *Endpoint) OutputGoIdent() protogen.GoIdent {
	return e.protoMethod.Output.GoIdent
}

func (e *Endpoint) ServerStreamName() string {
	method := e.protoMethod
	return method.Parent.GoName + "_" + method.GoName + "Server"
}

// --------------------- Http Rule ---------------------

func (e *Endpoint) ParseHttpRule() error {
	httpRule := proto.GetExtension(e.protoMethod.Desc.Options(), annotations.E_Http)
	if httpRule == nil || httpRule == annotations.E_Http.InterfaceOf(annotations.E_Http.Zero()) {
		pattern := &annotations.HttpRule_Post{
			Post: e.FullName(),
		}
		body := "*"
		httpRule = &annotations.HttpRule{
			Pattern: pattern,
			Body:    body,
		}
	}
	e.httpRule = httpRule.(*annotations.HttpRule)
	if err := e.ParsePattern(); err != nil {
		return err
	}
	e.ParsePath()
	if err := e.ParseRequestBody(); err != nil {
		return err
	}
	if err := e.ParseNamedPathFields(); err != nil {
		return err
	}
	if err := e.ParsePathFields(); err != nil {
		return err
	}
	if err := e.ParseQueryFields(); err != nil {
		return err
	}
	return nil
}

func (e *Endpoint) ParsePattern() error {
	switch pattern := e.httpRule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		e.httpMethod = MethodGet
		e.rawPath = pattern.Get
		return nil
	case *annotations.HttpRule_Post:
		e.httpMethod = MethodPost
		e.rawPath = pattern.Post
		return nil
	case *annotations.HttpRule_Put:
		e.httpMethod = MethodPut
		e.rawPath = pattern.Put
		return nil
	case *annotations.HttpRule_Delete:
		e.httpMethod = MethodDelete
		e.rawPath = pattern.Delete
		return nil
	case *annotations.HttpRule_Patch:
		e.httpMethod = MethodPatch
		e.rawPath = pattern.Patch
		return nil
	default:
		return fmt.Errorf("%s, unsupported httpMethod %s", e.FullName(), pattern)
	}
}

func (e *Endpoint) HttpMethod() any {
	return e.httpMethod
}

func (e *Endpoint) RawPath() string {
	return e.rawPath
}

func (e *Endpoint) ParsePath() {
	path := e.rawPath
	path, name, template, parameters := e.parseNamedPath(path)
	samplePathParameters := e.parseSamplePath(path)
	e.routePath = path
	e.namedPathFieldName = name
	e.namedPathTemplate = template
	e.namedPathFieldsParameters = parameters
	e.pathFieldsParameters = slicex.Remove(samplePathParameters, parameters...)
}

func (e *Endpoint) parseNamedPath(path string) (string, string, string, []string) {
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

func (e *Endpoint) parseSamplePath(path string) []string {
	// Find simple path parameters like {id}
	var parameters []string
	if allMatches := samplePathPattern.FindAllStringSubmatch(path, -1); allMatches != nil {
		for _, matches := range allMatches {
			pathParameter := matches[1]
			parameters = append(parameters, pathParameter)
			path = strings.Replace(path, matches[1], pathParameter, 1)
		}
	}
	return parameters
}

func (e *Endpoint) RoutePath() string {
	return e.routePath
}

func (e *Endpoint) NamedPathFieldName() string {
	return e.namedPathFieldName
}

func (e *Endpoint) NamedPathTemplate() string {
	return e.namedPathTemplate
}

func (e *Endpoint) NamedPathFieldsParameters() []string {
	return e.namedPathFieldsParameters
}

func (e *Endpoint) PathFieldsParameters() []string {
	return e.pathFieldsParameters
}

func (e *Endpoint) ParseRequestBody() error {
	// body arguments
	var bodyField *protogen.Field
	body := e.httpRule.GetBody()
	switch body {
	case "":
		// ignore
		return nil
	case "*":
		// all
		e.bodyMessage = e.Input()
		return nil
	default:
		// field
		bodyField = FindField(body, e.Input())
		if bodyField == nil {
			return fmt.Errorf("%s, failed to find body field %s", e.FullName(), body)
		}
		e.bodyField = bodyField
		return nil
	}
}

func (e *Endpoint) BodyMessage() *protogen.Message {
	return e.bodyMessage
}

func (e *Endpoint) BodyField() *protogen.Field {
	return e.bodyField
}

func (e *Endpoint) ParseNamedPathFields() error {
	var namedPathFields []*protogen.Field
	message := e.Input()
	// 根据path解析出来的NamedPathFieldName查找
	if namedPathName := e.NamedPathFieldName(); len(namedPathName) > 0 {
		// 以.分割
		namedPathParameters := strings.Split(namedPathName, ".")
		// 逐级查找
		for i, namedPathParameter := range namedPathParameters {
			field := FindField(namedPathParameter, message)
			// 如果找不到，返回错误
			if field == nil {
				return fmt.Errorf("%s, failed to find named path parameter, %s", e.FullName(), namedPathName)
			}

			// 如果不是最后一级参数，那么必须是message类型
			if i < len(namedPathParameters)-1 {
				// 如果不是message类型，返回错误
				if field.Desc.Kind() != protoreflect.MessageKind {
					return fmt.Errorf("%s, named path parameter '%s' is not message, %s", e.FullName(), field.Desc.Name(), namedPathName)
				}
			} else {
				// 如果是最后一级参数，那么必须是string或者google.protobuf.StringValue类型
				switch field.Desc.Kind() {
				case protoreflect.StringKind:
				case protoreflect.MessageKind:
					switch field.Message.Desc.FullName() {
					case "google.protobuf.StringValue":
					default:
						return fmt.Errorf("%s, named path parameter only support 'string' or 'google.protobuf.StringValue' type, %s", e.FullName(), namedPathName)
					}
				default:
					return fmt.Errorf("%s, named path parameter only support 'string' or 'google.protobuf.StringValue' type, %s", e.FullName(), namedPathName)
				}
			}
			// 下一级查找
			namedPathFields = append(namedPathFields, field)
			message = field.Message
		}
	}
	e.namedPathFields = namedPathFields
	return nil
}

func (e *Endpoint) NamedPathFields() []*protogen.Field {
	return e.namedPathFields
}

func (e *Endpoint) ParsePathFields() error {
	var pathFields []*protogen.Field
	for _, pathParameter := range e.PathFieldsParameters() {
		field := FindField(pathParameter, e.Input())
		if field == nil {
			return fmt.Errorf("%s, failed to find path parameter, %s", e.FullName(), pathParameter)
		}
		if field.Desc.IsList() || field.Desc.IsMap() {
			return fmt.Errorf("%s, path parameter do not support list or map, %s", e.FullName(), pathParameter)
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
				return fmt.Errorf("%s, path parameters do not support %s", e.FullName(), message.Desc.FullName())
			}
		default:
			return fmt.Errorf("%s, path parameters do not support %s", e.FullName(), field.Desc.Kind())
		}
		pathFields = append(pathFields, field)
	}
	e.pathFields = pathFields
	return nil
}

func (e *Endpoint) PathFields() []*protogen.Field {
	return e.pathFields
}

func (e *Endpoint) ParseQueryFields() error {
	var queryFields []*protogen.Field
	if e.BodyMessage() != nil {
		// 如果bodyMessage不为空，那么忽略query参数
		return nil
	}
	for _, field := range e.Input().Fields {
		if field == e.BodyField() {
			// 跳过bodyField
			continue
		}
		if slices.Contains(e.NamedPathFields(), field) {
			// 跳过namedPathFields
			continue
		}
		if slices.Contains(e.PathFields(), field) {
			// 跳过pathFields
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
	e.queryFields = queryFields
	return nil
}

func (e *Endpoint) QueryFields() []*protogen.Field {
	return e.queryFields
}

func (e *Endpoint) RequestBody() string {
	return e.httpRule.GetBody()
}

func (e *Endpoint) ResponseBody() string {
	return e.httpRule.GetResponseBody()
}

func (e *Endpoint) HandlerName() string {
	return e.service.ServiceName() + "_" + e.Name() + "_Handler"
}

func (e *Endpoint) HandlerVarName() string {
	return e.Name() + "Handler"
}

// --------------------- Cqrs ---------------------

func (e *Endpoint) IsCommand() bool {
	return len(e.Output().Fields)+len(e.Output().Oneofs) == 0
}

func (e *Endpoint) IsQuery() bool {
	return !e.IsCommand()
}

func (e *Endpoint) CommandName() string {
	return e.Name() + "Command"
}
func (e *Endpoint) QueryName() string {
	return e.Name() + "Query"
}
func (e *Endpoint) ResultName() string {
	return e.Name() + "Result"
}

func (e *Endpoint) CommandTypeName() string {
	return e.CommandName() + "Type"
}
func (e *Endpoint) QueryTypeName() string {
	return e.QueryName() + "Type"
}
func (e *Endpoint) ResultTypeName() string {
	return e.ResultName() + "Type"
}

func (e *Endpoint) IsCommandMethod() string {
	return "is" + e.CommandName() + "_Kind"
}
func (e *Endpoint) IsQueryMethod() string {
	return "is" + e.QueryName() + "_Kind"
}
func (e *Endpoint) IsResultMethod() string {
	return "is" + e.ResultName() + "_Kind"
}

func (e *Endpoint) UnimplementedCommandName() string {
	return "Unimplemented" + e.CommandName()
}
func (e *Endpoint) UnimplementedQueryName() string {
	return "Unimplemented" + e.QueryName()
}
func (e *Endpoint) UnimplementedResultName() string {
	return "Unimplemented" + e.ResultName()
}
