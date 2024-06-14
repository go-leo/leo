package sdx

import (
	"context"
	"github.com/go-leo/gox/slicex"
	"golang.org/x/exp/slices"
	"strings"
)

type Colors []*Color

type Color struct {
	service string
	values  []string
}

// Pairs converts the colors to a slice of strings.
func (colors Colors) Pairs() []string {
	pairs := make([]string, 0, len(colors))
	for _, c := range colors {
		pairs = append(pairs, c.service+"="+strings.Join(c.values, ","))
	}
	return pairs
}

func (colors Colors) Find(service string) (*Color, bool) {
	return slicex.FindFunc(colors, func(color *Color) bool { return color.service == service })
}

func (color *Color) Color() []string {
	if color == nil {
		return nil
	}
	cloned := slices.Clone(color.values)
	slices.Sort(cloned)
	return cloned
}

// ParseColors parses the colors from the string.
// pair like 'google.example.library=red,blue,green'
// All target URLs like 'consul://.../...' will be resolved by this builder
func ParseColors(pairs []string) Colors {
	colors := make(Colors, 0, len(pairs))
	for _, value := range pairs {
		pair := strings.Split(value, "=")
		if len(pair) != 2 {
			continue
		}
		colorValues := strings.Split(pair[1], ",")
		colors = append(colors, &Color{service: pair[0], values: colorValues})
	}
	return colors
}

type colorKey struct{}

// InjectColors injects the colors into the context.
func InjectColors(ctx context.Context, colors Colors) context.Context {
	return context.WithValue(ctx, colorKey{}, colors)
}

// ExtractColors extracts the colors from the context.
func ExtractColors(ctx context.Context) (Colors, bool) {
	v, ok := ctx.Value(colorKey{}).(Colors)
	return v, ok
}
