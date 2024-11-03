package betterReflect

import (
	"fmt"
	"reflect"
	"strconv"
)

// SetValueFromString parses the string to the type of the value
// and sets it. Supported types are
//
// int, int8, int16, int32, int64, float32, float63, bool, string
func SetValueFromString(elem reflect.Value, s string) error {
	switch elem.Kind() {
	case reflect.String:
		elem.SetString(s)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return fmt.Errorf("cannot convert %q to INT", s)
		}
		elem.SetInt(i)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return fmt.Errorf("cannot convert %q to FLOAT", s)
		}
		elem.SetFloat(f)
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return fmt.Errorf("cannot convert %q to BOOL", s)
		}
		elem.SetBool(b)
	default:
		return fmt.Errorf("cannot set type %q", elem.Kind())
	}
	return nil
}
