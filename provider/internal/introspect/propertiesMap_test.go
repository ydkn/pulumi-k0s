// Copyright 2025, Florian Schwab.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package introspect

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

const stringPtr = "string"

func TestGetPlainPropertiesMap(t *testing.T) {
	stringPtr := stringPtr
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
		String:    stringPtr,
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
		stringPtr, "stringPtr",
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
