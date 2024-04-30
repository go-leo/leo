package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

// Comments is a comments string as provided by protoc.
type Comments string

// String formats the comments by inserting // to the start of each line,
// ensuring that there is a trailing newline.
// An empty comment is formatted as an empty string.
func (c Comments) String() string {
	if c == "" {
		return ""
	}
	var b []byte
	for _, line := range strings.Split(strings.TrimSuffix(string(c), "\n"), "\n") {
		b = append(b, "//"...)
		b = append(b, line...)
	}
	return string(b)
}

// singular produces the singular form of a collection name.
func singular(plural string) string {
	if strings.HasSuffix(plural, "ves") {
		return strings.TrimSuffix(plural, "ves") + "f"
	}
	if strings.HasSuffix(plural, "ies") {
		return strings.TrimSuffix(plural, "ies") + "y"
	}
	if strings.HasSuffix(plural, "s") {
		return strings.TrimSuffix(plural, "s")
	}
	return plural
}

func FindField(name string, inMessage *protogen.Message) *protogen.Field {
	for _, field := range inMessage.Fields {
		if string(field.Desc.Name()) == name || field.Desc.JSONName() == name {
			return field
		}
	}
	return nil
}

// FullMessageTypeName builds the full type name of a message.
func FullMessageTypeName(message protoreflect.MessageDescriptor) string {
	name := GetMessageName(message)
	return "." + string(message.ParentFile().Package()) + "." + name
}

func GetMessageName(message protoreflect.MessageDescriptor) string {
	prefix := ""
	parent := message.Parent()

	if _, ok := parent.(protoreflect.MessageDescriptor); ok {
		prefix = string(parent.Name()) + "_" + prefix
	}
	return prefix + string(message.Name())
}

func FullFieldName(fields []*protogen.Field) string {
	var fieldNames []string
	for _, p := range fields {
		fieldNames = append(fieldNames, p.GoName)
	}
	fullFieldName := strings.Join(fieldNames, ".")
	return fullFieldName
}
