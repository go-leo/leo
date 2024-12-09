package yamlx

import (
	"github.com/go-leo/gox/encodingx/jsonx"
	"github.com/go-leo/gox/encodingx/yamlx"
	"github.com/go-leo/leo/v3/configx"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

var _ configx.Parser = (*Parser)(nil)

type Parser struct{}

func (p *Parser) Support(format configx.Formatter) bool {
	return slices.ContainsFunc([]string{"yaml", "yml"}, func(s string) bool {
		return strings.EqualFold(format.Format(), s)
	})
}

func (p *Parser) Parse(source []byte) (*structpb.Struct, error) {
	v := make(map[string]any)
	if err := yamlx.Unmarshal(source, &v); err != nil {
		return nil, err
	}
	jsonData, err := jsonx.Marshal(v)
	if err != nil {
		return nil, err
	}
	value := &structpb.Struct{Fields: make(map[string]*structpb.Value)}
	return value, value.UnmarshalJSON(jsonData)
}
