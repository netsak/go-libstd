package betterReflect_test

import (
	"reflect"
	"testing"

	betterReflect "github.com/netsak/go-libstd/reflect"
)

func TestSetValueFromString_String(t *testing.T) {
	var got string = ""
	var want string = "string"
	betterReflect.SetValueFromString(reflect.ValueOf(&got).Elem(), "string")
	if want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestSetValueFromString_Int(t *testing.T) {
	var got int = 0
	var want int = 1909
	betterReflect.SetValueFromString(reflect.ValueOf(&got).Elem(), "1909")
	if want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
	err := betterReflect.SetValueFromString(reflect.ValueOf(&got).Elem(), "krosch")
	wantErr := `cannot convert "krosch" to INT`
	if err.Error() != wantErr {
		t.Errorf("want: %v, got: %v", wantErr, err.Error())
	}
}

func TestSetValueFromString_Float(t *testing.T) {
	var got float64 = 0
	var want float64 = 19.09
	betterReflect.SetValueFromString(reflect.ValueOf(&got).Elem(), "19.09")
	if want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
	err := betterReflect.SetValueFromString(reflect.ValueOf(&got).Elem(), "krosch")
	wantErr := `cannot convert "krosch" to FLOAT`
	if err.Error() != wantErr {
		t.Errorf("want: %v, got: %v", wantErr, err.Error())
	}
}

func TestSetValueFromString_Bool(t *testing.T) {
	var got bool = true
	var want bool = true
	betterReflect.SetValueFromString(reflect.ValueOf(&got).Elem(), "true")
	if want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
	err := betterReflect.SetValueFromString(reflect.ValueOf(&got).Elem(), "krosch")
	wantErr := `cannot convert "krosch" to BOOL`
	if err.Error() != wantErr {
		t.Errorf("want: %v, got: %v", wantErr, err.Error())
	}
}

func TestSetValueFromString_UnknownType(t *testing.T) {
	var got interface{}
	err := betterReflect.SetValueFromString(reflect.ValueOf(&got).Elem(), "true")
	wantErr := `cannot set type "interface"`
	if err.Error() != wantErr {
		t.Errorf("want: %v, got: %v", wantErr, err.Error())
	}
}
