package conf

import (
	"context"
	"google.golang.org/protobuf/types/known/structpb"
)

type ApplicationConfig interface {
	GetConfig() *Application
	GetValue(key string) *structpb.Value
	Refresh(ctx context.Context) error
	Notify(ctx context.Context, ch chan<- *Application) error
	StopNotify(ch chan<- *Application) error
}

type Source struct {
	Name string
	Data []byte
}

type Format string

// Resource is a loader that can be used to load source config.
type Resource interface {
	Load(ctx context.Context) (Source, error)
	Notify(ctx context.Context, ch chan<- Source)
	StopNotify(ch chan<- Source)
	Format() Format
}

type Parser interface {
	Parse(source Source) (*structpb.Struct, error)
	Support(format Format) bool
}

type Merger interface {
	Merge(targets ...*structpb.Struct) *structpb.Struct
}

type applicationConfig struct {
	resource []Resource
	parser   []Parser
}

func (c *applicationConfig) GetConfig() *Application {
	newStruct, err := structpb.NewStruct()
	s, err := structpb.NewStruct()

	//TODO implement me
	panic("implement me")
}

func (c *applicationConfig) GetValue(key string) *structpb.Value {
	//TODO implement me
	panic("implement me")
}

func (c *applicationConfig) Refresh(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *applicationConfig) Notify(ch chan<- *Application) error {
	//TODO implement me
	panic("implement me")
}

func (c *applicationConfig) StopNotify(ch chan<- *Application) error {
	//TODO implement me
	panic("implement me")
}

func newApplicationConfig(ctx context.Context, resources []Resource, parsers []Parser, merger Merger) (ApplicationConfig, error) {
	var targets []*structpb.Struct
	for _, resource := range resources {
		for _, parser := range parsers {
			if parser.Support(resource.Format()) {
				source, err := resource.Load(ctx)
				if err != nil {
					return nil, err
				}
				target, err := parser.Parse(source)
				if err != nil {
					return nil, err
				}
				targets = append(targets, target)
			}
		}
	}
	target := merger.Merge(targets...)
	for key, value := range target.GetFields() {
		kind := value.GetKind()
		value.AsInterface()
	}
	return &applicationConfig{}
}
