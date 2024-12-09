package tomlx

import (
	"github.com/go-leo/gox/encodingx/jsonx"
	"github.com/go-leo/gox/encodingx/tomlx"
	"github.com/go-leo/leo/v3/configx"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

var _ configx.Parser = (*Parser)(nil)

type Parser struct{}

func (p *Parser) Support(format configx.Formatter) bool {
	return strings.EqualFold(format.Format(), "toml")
}

func (p *Parser) Parse(source []byte) (*structpb.Struct, error) {
	v := make(map[string]any)
	if err := tomlx.Unmarshal(source, &v); err != nil {
		return nil, err
	}
	jsonData, err := jsonx.Marshal(v)
	if err != nil {
		return nil, err
	}
	value := &structpb.Struct{Fields: make(map[string]*structpb.Value)}
	return value, value.UnmarshalJSON(jsonData)
}
