package configx

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/configx/test"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
	"time"
)

// TestWatch tests the Watch function
func TestWatch(t *testing.T) {
	value, _ := structpb.NewStruct(map[string]any{
		"addr": "127.0.0.1",
		"port": 8080,
	})
	data, _ := value.MarshalJSON()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var notifyC chan<- *Event

	mockParser := new(MockParser)
	mockResource := new(MockResource)

	mockParser.On("Support", mockResource).Return(true)
	mockParser.On("Parse", data).Return(value, nil)

	mockResource.On("Format", mock.Anything).Return("json")
	mockResource.On("Load", ctx).Return(data, nil)
	mockResource.On("Watch", ctx, mock.Anything).Run(func(args mock.Arguments) {
		notifyC = args.Get(1).(chan<- *Event)
		go func() {
			for i := 0; i < 10; i++ {
				if randx.Bool() {
					notifyC <- &Event{kind: &DataEvent{Data: nil}}
				} else {
					notifyC <- &Event{kind: &ErrorEvent{Err: fmt.Errorf("error")}}
				}
				<-time.After(time.Second)
			}
		}()
	}).Return(func() {
		notifyC <- &Event{kind: &ErrorEvent{Err: ErrStopWatch}}
	}, nil)

	confC, errC, stop := Watch[*test.Application](ctx, WithResource(mockResource), WithParser(mockParser))

	go func() {
		<-time.After(12 * time.Second)
		stop()
	}()

loop:
	for {
		select {
		case conf := <-confC:
			fmt.Println("conf:", conf)
		case err := <-errC:
			if errors.Is(err, ErrStopWatch) {
				break loop
			}
			fmt.Println("err:", err)
		}
	}

}
