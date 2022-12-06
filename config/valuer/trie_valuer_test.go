package valuer_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-leo/leo/v2/config/valuer"
)

func TestTrieTreeValuer(t *testing.T) {

	key1 := "val1"
	key2 := 1
	key3 := true
	key4 := time.Now()
	key5 := time.Second
	key6 := 1.414
	subconf2key2 := []string{"a", "b", "c"}
	subconf2key3 := []int{1, 2, 3, 4, 5}
	subconf2subconf3key2 := "val2"

	subconfkey1 := "val1"
	subconf := map[string]any{
		"key1": subconfkey1,
	}

	subconf2subconf3key1 := "val1"
	subconf2subconf3 := map[string]any{
		"key1": subconf2subconf3key1,
		"key2": subconf2subconf3key2,
	}
	subconf2 := map[any]any{
		"key2":     subconf2key2,
		"key3":     subconf2key3,
		"subconf3": subconf2subconf3,
	}
	config1 := map[string]any{
		"key1":     key1,
		"key2":     key2,
		"key3":     key3,
		"key4":     key4,
		"key5":     key5,
		"key6":     key6,
		"subconf":  subconf,
		"subconf2": subconf2,
	}

	newKey6 := 3.1415
	newKey3 := false
	newsubconf2key2 := []string{"d", "e"}
	config2 := map[string]any{
		"key3": newKey3,
		"key6": newKey6,
		"subconf2": map[any]any{
			"key2": newsubconf2key2,
		},
	}

	valuer := valuer.NewTrieTreeValuer()
	valuer.AddConfig(config1, config2)

	key1Val, err := valuer.Get("key1")
	if err != nil {
		t.Fatal(err)
	}
	if key1Val != key1 {
		t.Fatalf("expected value is %s,but actual is %s", key1, key1Val)
	}

	key6Val, err := valuer.Get("key6")
	if err != nil {
		t.Fatal(err)
	}
	if key6Val != newKey6 {
		t.Fatalf("expected value is %f,but actual is %f", key6, key6Val)
	}

	key3Val, err := valuer.Get("key3")
	if err != nil {
		t.Fatal(err)
	}
	if key3Val != newKey3 {
		t.Fatalf("expected value is %v,but actual is %v", newKey3, key3Val)
	}
	subconf2key2Val, err := valuer.Get("subconf2.key2")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(subconf2key2Val, newsubconf2key2) {
		t.Fatalf("expected value is %v,but actual is %v", newsubconf2key2, subconf2key2Val)
	}

	subconfVal, err := valuer.Get("subconf")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(subconfVal, subconf) {
		t.Fatalf("expected value is %v,but actual is %v", subconf, subconfVal)
	}
}
