package config

import (
	"cmp"
	"fmt"
	"os"
	"reflect"
	"strings"

	betterReflect "github.com/netsak/go-libstd/reflect"
	"github.com/netsak/go-libstd/text"
)

func New[T any](conf *T) (*T, error) {
	elem := reflect.ValueOf(conf).Elem()
	typ := reflect.TypeOf(conf).Elem()
	if err := set(elem, typ, ""); err != nil {
		return nil, err
	}
	return conf, nil
}

func set(elem reflect.Value, typ reflect.Type, prefix string) error {
	if prefix != "" && prefix != "_" {
		prefix += "_"
	}
	for i := 0; i < elem.NumField(); i++ {
		fieldElem := elem.Field(i)
		if !fieldElem.CanSet() {
			continue // cannot set the field
		}
		fieldType := typ.Field(i)
		key := cmp.Or(fieldType.Tag.Get("env"), text.ToEnvVar(fieldType.Name))
		key = strings.ToUpper(prefix + key)
		if fieldElem.Kind() == reflect.Struct { // recurse into struct
			set(fieldElem, fieldType.Type, key)
			continue
		}
		value, ok := os.LookupEnv(key)
		if !ok {
			continue // no value found
		}
		if err := betterReflect.SetValueFromString(fieldElem, value); err != nil {
			return fmt.Errorf("field %s: %w", fieldType.Name, err)
		}
	}
	return nil
}

func MustNew[T any](conf *T) *T {
	conf, err := New[T](conf)
	if err != nil {
		panic(err)
	}
	return conf
}
