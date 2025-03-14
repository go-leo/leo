package metadatax

import (
	"context"
	"reflect"
	"strconv"
	"testing"
	"time"
)

const defaultTestTimeout = 10 * time.Second

func TestPairsMD(t *testing.T) {
	for _, test := range []struct {
		// input
		kv []string
		// output
		md Metadata
	}{
		{[]string{}, New()},
		{[]string{"k1", "v1", "k1", "v2"}, FromMap(map[string][]string{"k1": {"v1", "v2"}})},
	} {
		md := Pairs(test.kv...)
		if !reflect.DeepEqual(md, test.md) {
			t.Fatalf("Pairs(%v) = %v, want %v", test.kv, md, test.md)
		}
	}
}

func TestCopy(t *testing.T) {
	const key, val = "key", "val"
	orig := New()
	orig.Set(key, val)
	cpy := orig.Clone()
	if !reflect.DeepEqual(orig, cpy) {
		t.Errorf("copied value not equal to the original, got %v, want %v", cpy, orig)
	}
	origM := orig.(_Metadata)
	cpyM := cpy.(_Metadata)
	origM[key][0] = "foo"
	if v := cpyM[key][0]; v != val {
		t.Errorf("change in original should not affect copy, got %q, want %q", v, val)
	}
}

func TestJoin(t *testing.T) {
	for _, test := range []struct {
		mds  []Metadata
		want Metadata
	}{
		{[]Metadata{}, New()},
		{[]Metadata{Pairs("foo", "bar")}, Pairs("foo", "bar")},
		{[]Metadata{Pairs("foo", "bar"), Pairs("foo", "baz")}, Pairs("foo", "bar", "foo", "baz")},
		{[]Metadata{Pairs("foo", "bar"), Pairs("foo", "baz"), Pairs("zip", "zap")}, Pairs("foo", "bar", "foo", "baz", "zip", "zap")},
	} {
		md := Join(test.mds...)
		if !reflect.DeepEqual(md, test.want) {
			t.Errorf("context's metadata is %v, want %v", md, test.want)
		}
	}
}

func TestValues(t *testing.T) {
	for _, test := range []struct {
		md       Metadata
		key      string
		wantVals []string
	}{
		{md: Pairs("My-Optional-Header", "42"), key: "My-Optional-Header", wantVals: []string{"42"}},
		{md: Pairs("Header", "42", "Header", "43", "Header", "44", "other", "1"), key: "HEADER", wantVals: []string{"42", "43", "44"}},
		{md: Pairs("HEADER", "10"), key: "HEADER", wantVals: []string{"10"}},
	} {
		vals := test.md.Values(test.key)
		if !reflect.DeepEqual(vals, test.wantVals) {
			t.Errorf("value of metadata %v is %v, want %v", test.key, vals, test.wantVals)
		}
	}
}

func TestSet(t *testing.T) {
	for _, test := range []struct {
		md      Metadata
		setKey  string
		setVals []string
		want    Metadata
	}{
		{
			md:      Pairs("My-Optional-Header", "42", "other-key", "999"),
			setKey:  "Other-Key",
			setVals: []string{"1"},
			want:    Pairs("my-optional-header", "42", "other-key", "1"),
		},
		{
			md:      Pairs("My-Optional-Header", "42"),
			setKey:  "Other-Key",
			setVals: []string{"1", "2", "3"},
			want:    Pairs("my-optional-header", "42", "other-key", "1", "other-key", "2", "other-key", "3"),
		},
		{
			md:      Pairs("My-Optional-Header", "42"),
			setKey:  "Other-Key",
			setVals: []string{},
			want:    Pairs("my-optional-header", "42"),
		},
	} {
		test.md.Set(test.setKey, test.setVals...)
		if !reflect.DeepEqual(test.md, test.want) {
			t.Errorf("value of metadata is %v, want %v", test.md, test.want)
		}
	}
}

func TestAppend(t *testing.T) {
	for _, test := range []struct {
		md         Metadata
		appendKey  string
		appendVals []string
		want       Metadata
	}{
		{
			md:         Pairs("My-Optional-Header", "42"),
			appendKey:  "Other-Key",
			appendVals: []string{"1"},
			want:       Pairs("my-optional-header", "42", "other-key", "1"),
		},
		{
			md:         Pairs("My-Optional-Header", "42"),
			appendKey:  "my-OptIoNal-HeAder",
			appendVals: []string{"1", "2", "3"},
			want: Pairs("my-optional-header", "42", "my-optional-header", "1",
				"my-optional-header", "2", "my-optional-header", "3"),
		},
		{
			md:         Pairs("My-Optional-Header", "42"),
			appendKey:  "my-OptIoNal-HeAder",
			appendVals: []string{},
			want:       Pairs("my-optional-header", "42"),
		},
	} {
		test.md.Append(test.appendKey, test.appendVals...)
		if !reflect.DeepEqual(test.md, test.want) {
			t.Errorf("value of metadata is %v, want %v", test.md, test.want)
		}
	}
}

func TestDelete(t *testing.T) {
	for _, test := range []struct {
		md        Metadata
		deleteKey string
		want      Metadata
	}{
		{
			md:        Pairs("My-Optional-Header", "42"),
			deleteKey: "My-Optional-Header",
			want:      Pairs(),
		},
		{
			md:        Pairs("My-Optional-Header", "42"),
			deleteKey: "Other-Key",
			want:      Pairs("my-optional-header", "42"),
		},
		{
			md:        Pairs("My-Optional-Header", "42"),
			deleteKey: "my-OptIoNal-HeAder",
			want:      Pairs(),
		},
	} {
		test.md.Delete(test.deleteKey)
		if !reflect.DeepEqual(test.md, test.want) {
			t.Errorf("value of metadata is %v, want %v", test.md, test.want)
		}
	}
}

//func TestFromIncomingContext(t *testing.T) {
//	_Metadata := Pairs("X-My-Header-1", "42").(_Metadata)
//	// Verify that we lowercase if callers directly modify _Metadata
//	_Metadata["X-INCORRECT-UPPERCASE"] = []string{"foo"}
//	ctx := NewIncomingContext(context.Background(), _Metadata)
//
//	result, found := FromIncomingContext(ctx)
//	if !found {
//		t.Fatal("FromIncomingContext must return metadata")
//	}
//	resultM := result.(_Metadata)
//	expected := _Metadata{
//		"x-my-header-1":         []string{"42"},
//		"x-incorrect-uppercase": []string{"foo"},
//	}
//	if !reflect.DeepEqual(result, expected) {
//		t.Errorf("FromIncomingContext returned %#v, expected %#v", result, expected)
//	}
//
//	// ensure modifying result does not modify the value in the context
//	resultM["new_key"] = []string{"foo"}
//	resultM["x-my-header-1"][0] = "mutated"
//
//	result2, found := FromIncomingContext(ctx)
//	if !found {
//		t.Fatal("FromIncomingContext must return metadata")
//	}
//	if !reflect.DeepEqual(result2, expected) {
//		t.Errorf("FromIncomingContext after modifications returned %#v, expected %#v", result2, expected)
//	}
//}

func TestAppendToOutgoingContext(t *testing.T) {
	// Pre-existing metadata
	tCtx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()
	ctx := NewOutgoingContext(tCtx, Pairs("k1", "v1", "k2", "v2"))
	ctx = AppendOutgoingContext(ctx, Pairs("k1", "v3"))
	ctx = AppendOutgoingContext(ctx, Pairs("k1", "v4"))
	md, ok := FromOutgoingContext(ctx)
	if !ok {
		t.Errorf("Expected MD to exist in ctx, but got none")
	}
	want := Pairs("k1", "v1", "k1", "v3", "k1", "v4", "k2", "v2")
	if !reflect.DeepEqual(md, want) {
		t.Errorf("context's metadata is %v, want %v", md, want)
	}

	// No existing metadata
	ctx = AppendOutgoingContext(tCtx, Pairs("k1", "v1"))
	md, ok = FromOutgoingContext(ctx)
	if !ok {
		t.Errorf("Expected MD to exist in ctx, but got none")
	}
	want = Pairs("k1", "v1")
	if !reflect.DeepEqual(md, want) {
		t.Errorf("context's metadata is %v, want %v", md, want)
	}
}

func TestAppendToOutgoingContext_Repeated(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()

	for i := 0; i < 100; i = i + 2 {
		ctx1 := AppendOutgoingContext(ctx, Pairs("k", strconv.Itoa(i)))
		ctx2 := AppendOutgoingContext(ctx, Pairs("k", strconv.Itoa(i+1)))

		md1, _ := FromOutgoingContext(ctx1)
		md2, _ := FromOutgoingContext(ctx2)

		if reflect.DeepEqual(md1, md2) {
			t.Fatalf("md1, md2 = %v, %v; should not be equal", md1, md2)
		}

		ctx = ctx1
	}
}

//func TestAppendToOutgoingContext_FromKVSlice(t *testing.T) {
//	const k, v = "a", "b"
//	kv := []string{k, v}
//	tCtx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
//	defer cancel()
//	ctx := AppendOutgoingContext(tCtx, Pairs(kv...))
//	_Metadata, _ := FromOutgoingContext(ctx)
//	mdM := _Metadata.(_Metadata)
//	if mdM[k][0] != v {
//		t.Fatalf("_Metadata[%q] = %q; want %q", k, mdM[k], v)
//	}
//	kv[1] = "xxx"
//	_Metadata, _ = FromOutgoingContext(ctx)
//	if mdM[k][0] != v {
//		t.Fatalf("_Metadata[%q] = %q; want %q", k, mdM[k], v)
//	}
//}
