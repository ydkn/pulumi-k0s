package introspect

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func TestGetPlainPropertiesMap(t *testing.T) {
	stringPtr := "string"
	boolPtr := true
	int64Ptr := int64(42)

	myStruct := struct {
		String    string  `pulumi:"string"`
		StringPtr *string `pulumi:"stringPtr"`
		Bool      bool    `pulumi:"bool"`
		BoolPtr   *bool   `pulumi:"boolPtr"`
		Int64     int64   `pulumi:"int64"`
		Int64Ptr  *int64  `pulumi:"int64Ptr"`
		NilPtr    *string `pulumi:"nilPtr"`
		Struct    struct {
			Foo string `pulumi:"foo"`
		} `pulumi:"struct"`
		StructPtr struct {
			Bar *string `pulumi:"bar"`
		} `pulumi:"structPtr"`
		Slice          []string  `pulumi:"slice"`
		SlicePtr       []*string `pulumi:"slicePtr"`
		SliceStructPtr []struct {
			Foo string `pulumi:"foo"`
			Bar string `pulumi:"bar"`
		} `pulumi:"sliceStructPtr"`
		Map    map[string]string  `pulumi:"map"`
		MapPtr map[string]*string `pulumi:"mapPtr"`
	}{
		String:    "string",
		StringPtr: &stringPtr,
		Bool:      true,
		BoolPtr:   &boolPtr,
		Int64:     42,
		Int64Ptr:  &int64Ptr,
		Struct: struct {
			Foo string `pulumi:"foo"`
		}{Foo: "foo"},
		StructPtr: struct {
			Bar *string `pulumi:"bar"`
		}{Bar: &stringPtr},
		Slice:    []string{"a", "b", "c"},
		SlicePtr: []*string{&stringPtr},
		SliceStructPtr: []struct {
			Foo string `pulumi:"foo"`
			Bar string `pulumi:"bar"`
		}{{Foo: "foo", Bar: "bar"}},
		Map:    map[string]string{"a": "b"},
		MapPtr: map[string]*string{"b": &stringPtr},
	}

	props, err := NewPropertiesMap(myStruct)
	if err != nil {
		t.Errorf("getPlainPropertiesMap error: %s", err.Error())

		return
	}

	keys := []string{
		"string", "stringPtr",
		"bool", "boolPtr",
		"int64", "int64Ptr",
		"struct.foo", "structPtr.bar",
		"slice[0]", "slice[1]", "slice[2]", "slicePtr[0]", "sliceStructPtr[0].foo", "sliceStructPtr[0].bar",
		"nilPtr",
		"map.a", "mapPtr.b",
	}

	keysMap := map[string]any{}
	for _, k := range keys {
		keysMap[k] = nil
	}

	for k := range props {
		if _, ok := keysMap[string(k)]; !ok {
			t.Errorf("unexpected key: %s", k)
		}
	}

	for _, k := range keys {
		if _, ok := props[resource.PropertyKey(k)]; !ok {
			t.Errorf("missing key: %s", k)
		}
	}
}
