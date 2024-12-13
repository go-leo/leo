package configx

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/configx/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
	"time"
)

var _ Resource = (*MockResource)(nil)

// MockResource is a mock struct that implements Resource interface
type MockResource struct {
	mock.Mock
}

func (m *MockResource) Load(ctx context.Context) ([]byte, error) {
	args := m.Called(ctx)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockResource) Format() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *MockResource) Watch(ctx context.Context, notifyC chan<- *Event) (func(), error) {
	args := m.Called(ctx, notifyC)

	return args.Get(0).(func()), args.Error(1)
}

var _ Parser = (*MockParser)(nil)

type MockParser struct {
	mock.Mock
}

func (m *MockParser) Support(format Formatter) bool {
	args := m.Called(format)
	return args.Bool(0)
}

func (m *MockParser) Parse(data []byte) (*structpb.Struct, error) {
	args := m.Called(data)
	return args.Get(0).(*structpb.Struct), args.Error(1)
}

func TestLoad(t *testing.T) {
	value, _ := structpb.NewStruct(map[string]any{
		"addr": "127.0.0.1",
		"port": 8080,
	})
	data, _ := value.MarshalJSON()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mockParser := new(MockParser)
	mockResource := new(MockResource)

	mockParser.On("Support", mockResource).Return(true)
	mockParser.On("Parse", data).Return(value, nil)

	mockResource.On("Format", mock.Anything).Return("json")
	mockResource.On("Load", ctx).Return(data, nil)

	conf, err := Load[*test.Application](ctx, WithResource(mockResource), WithParser(mockParser))
	assert.Nil(t, err)

	fmt.Println(conf)
	// Give some time for channels to send data or for the function to return
	time.Sleep(100 * time.Millisecond)

}
