package response

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
)

func TestAny(t *testing.T) {
	anyUser, err := anypb.New(&User{
		Name:    "jax",
		Email:   "jax@github.com",
		Phone:   "8888888888",
		Address: "china shanghai",
	})
	if err != nil {
		panic(err)
	}
	t.Log(anyUser)

	anyString, err := anypb.New(wrapperspb.String("jax"))
	if err != nil {
		panic(err)
	}
	t.Log(anyString)

	dst, err := anypb.UnmarshalNew(anyString, proto.UnmarshalOptions{})
	if err != nil {
		panic(err)
	}
	t.Log(dst)

	st, err := structpb.NewStruct(
		map[string]interface{}{
			"name":   "jax",
			"age":    10,
			"height": 2.34,
			"sex":    true,
			"address": map[string]any{
				"city":   "shanghai",
				"street": "shanghai street",
			},
		})
	if err != nil {
		panic(err)
	}
	anyStruct, err := anypb.New(st)
	if err != nil {
		panic(err)
	}
	t.Log(anyStruct)

	structDst, err := anypb.UnmarshalNew(anyStruct, proto.UnmarshalOptions{})
	if err != nil {
		panic(err)
	}
	newSt := structDst.(*structpb.Struct)
	for key, value := range newSt.GetFields() {
		t.Log(key)
		data, _ := json.Marshal(value)
		t.Log(string(data))
	}
}
