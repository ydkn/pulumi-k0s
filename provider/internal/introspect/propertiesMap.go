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
	"fmt"
	"maps"
	"reflect"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func NewPropertiesMap(obj any) (resource.PropertyMap, error) {
	typ := reflect.TypeOf(obj)
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, found %s (%T)", typ.Kind(), obj)
	}

	return getPlainPropertiesMapWithPrefix(obj, "")
}

func getPlainPropertiesMapWithPrefix(obj any, prefix string) (resource.PropertyMap, error) {
	if prefix != "" {
		prefix += "."
	}

	typ := reflect.TypeOf(obj)
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return resource.PropertyMap{}, nil
	}

	val := reflect.ValueOf(obj)

	properties := resource.PropertyMap{}
	for _, f := range reflect.VisibleFields(typ) {
		fieldPrefix := prefix
		if val.Kind() != reflect.Slice {
			pulumiTag, ok := f.Tag.Lookup("pulumi")
			if !ok {
				continue
			}

			fieldPrefix += strings.Split(pulumiTag, ",")[0]
		}

		value := reflect.Indirect(val).FieldByName(f.Name)
		if value.Kind() == reflect.Pointer {
			value = value.Elem()
		}

		props, err := getPropertyMapForValue(value, fieldPrefix)
		if err != nil {
			return nil, err
		}

		for k, v := range props {
			properties[k] = v
		}
	}

	return properties, nil
}

func getPropertyMapForValue(value reflect.Value, prefix string) (resource.PropertyMap, error) {
	key := resource.PropertyKey(prefix)
	properties := resource.PropertyMap{}

	switch value.Kind() {
	case reflect.Struct:
		props, err := getPlainPropertiesMapWithPrefix(value.Interface(), prefix)
		if err != nil {
			return nil, err
		}

		maps.Copy(properties, props)
	case reflect.Slice:
		slice := reflect.ValueOf(value.Interface())

		for i := 0; i < slice.Len(); i++ {
			slicePrefix := fmt.Sprintf("%s[%d]", prefix, i)

			sliceElement := slice.Index(i)
			if sliceElement.Kind() == reflect.Pointer {
				sliceElement = sliceElement.Elem()
			}

			var props resource.PropertyMap
			var err error

			if sliceElement.Kind() == reflect.Struct {
				props, err = getPlainPropertiesMapWithPrefix(sliceElement.Interface(), slicePrefix)
			} else {
				props, err = getPropertyMapForValue(sliceElement, slicePrefix)
			}

			if err != nil {
				return nil, err
			}

			maps.Copy(properties, props)
		}
	case reflect.Map:
		mapValue := reflect.ValueOf(value.Interface())

		for _, e := range mapValue.MapKeys() {
			mapElement := mapValue.MapIndex(e)

			mapPrefix := fmt.Sprintf("%s.%s", prefix, e.String())

			if mapElement.Kind() == reflect.Pointer {
				mapElement = mapElement.Elem()
			}

			var props resource.PropertyMap
			var err error

			if mapElement.Kind() == reflect.Struct {
				props, err = getPlainPropertiesMapWithPrefix(mapElement, mapPrefix)
			} else {
				props, err = getPropertyMapForValue(mapElement, mapPrefix)
			}

			if err != nil {
				return nil, err
			}

			maps.Copy(properties, props)
		}
	case reflect.Bool:
		properties[key] = resource.NewBoolProperty(value.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		properties[key] = resource.NewNumberProperty(float64(value.Int()))
	case reflect.Float32, reflect.Float64:
		properties[key] = resource.NewNumberProperty(float64(value.Float()))
	case reflect.String:
		properties[key] = resource.NewStringProperty(value.String())
	case reflect.Invalid: // nil
		properties[key] = resource.NewNullProperty()
	default:
		return nil, fmt.Errorf("unsupported data type: %s", value.Kind())
	}

	return properties, nil
}
