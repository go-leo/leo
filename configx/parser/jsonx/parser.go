package jsonx

import (
	"github.com/go-leo/leo/v3/configx"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

var _ configx.Parser = (*Parser)(nil)

type Parser struct{}

func (p *Parser) Support(format configx.Formatter) bool {
	return strings.EqualFold(format.Format(), "json")
}

func (p *Parser) Parse(source []byte) (*structpb.Struct, error) {
	value := &structpb.Struct{Fields: make(map[string]*structpb.Value)}
	return value, value.UnmarshalJSON(source)
}
