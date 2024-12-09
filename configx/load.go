package configx

import (
	"context"
	"errors"
	"github.com/go-leo/gox/reflectx"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	ErrUnsupportedFormat = errors.New("config: unsupported format")
)

// Load 函数用于加载配置数据，通过资源加载器加载数据，通过解析器解析数据，最后通过合并器合并多个解析结果，并返回最终的配置对象。
func Load[Config proto.Message](ctx context.Context, opts ...Option) (Config, error) {
	// 初始化配置选项。
	opt := newOptions()
	opt.Apply(opts...)

	// 初始化配置对象。
	var errs []error

	// 遍历资源列表，加载数据并解析数据。
	var values []*structpb.Struct
	for _, loader := range opt.Resources {
		// 使用资源加载器加载数据
		source, err := loader.Load(ctx)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// 使用解析器解析数据
		foundParser := false
		for _, parser := range opt.Parsers {
			// 判断解析器是否支持当前资源格式
			if parser.Support(loader) {
				foundParser = true
				value, err := parser.Parse(source)
				if err != nil {
					errs = append(errs, err)
					break
				}
				values = append(values, value)
				break
			}
		}
		// 如果没有匹配的解析器，则记录错误
		if !foundParser {
			errs = append(errs, ErrUnsupportedFormat)
		}
	}

	// 将解析结果合并为一个配置对象。
	value := opt.Merger.Merge(values...)

	// 将合并结果序列化为JSON。
	data, err := value.MarshalJSON()
	if err != nil {
		errs = append(errs, err)
	}

	// 将JSON反序列化为配置对象。
	conf := reflectx.New[Config]()
	var p proto.Message = conf
	if err := protojson.Unmarshal(data, p); err != nil {
		errs = append(errs, err)
	}
	return conf, errors.Join(errs...)
}
