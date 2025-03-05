package leo

import (
	"context"
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp()
	_ = app.Run(context.Background())
}

func TestWrapperspb(t *testing.T) {
	value := wrapperspb.Bool(true)
	data, err := protojson.Marshal(value)
	if err != nil {
		panic(err)
	}
	t.Log(string(data))
}

func TestStructpb(t *testing.T) {
	value := structpb.NewBoolValue(true)
	data, err := protojson.Marshal(value)
	if err != nil {
		panic(err)
	}
	t.Log(string(data))
}

func TestMultiBytes(t *testing.T) {
	data, err := json.Marshal([][]byte{{'a', 'b', 'c'}, {'1', '2', '3'}, {'+', '-', '*'}})
	if err != nil {
		panic(err)
	}
	t.Log(string(data))
}
