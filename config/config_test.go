package config_test

import (
	"reflect"
	"testing"

	"github.com/netsak/go-libstd/config"
)

type Sub struct {
	AccessToken string
	Name        string
	Explicit    string `env:"MY_EXPLICIT"`
}

type SomeStruct struct {
	String      string
	Tag         string `env:"BOB"`
	Nothing     string
	Default     string
	Overwritten string
	Int         int
	Float64     float64
	Float32     float32
	Bool        bool
	Sub         Sub
	private     string // not settable
}

func TestPlainStruct(t *testing.T) {
	t.Setenv("STRING", "my-string")
	t.Setenv("OVERWRITTEN", "overwritten")
	t.Setenv("BOB", "bob-env-var")
	t.Setenv("INT", "42")
	t.Setenv("FLOAT64", "19.09")
	t.Setenv("FLOAT32", "1.2")
	t.Setenv("BOOL", "true")
	t.Setenv("PRIVATE", "not-settable")
	t.Setenv("SUB_MY_EXPLICIT", "my-explicit")
	t.Setenv("SUB_NAME", "sub-name")
	t.Setenv("SUB_ACCESS_TOKEN", "access-token")
	conf := SomeStruct{
		Default:     "default",
		Overwritten: "original",
	}
	got, err := config.New(&conf)
	if err != nil {
		t.Errorf("expected no error, got: %s", err)
	}
	want := SomeStruct{
		String:      "my-string",
		Tag:         "bob-env-var",
		Nothing:     "",
		Default:     "default",
		Overwritten: "overwritten",
		Int:         42,
		Float64:     19.09,
		Float32:     1.2,
		Bool:        true,
		Sub: Sub{
			Name:        "sub-name",
			Explicit:    "my-explicit",
			AccessToken: "access-token",
		},
	}
	if !reflect.DeepEqual(conf, want) {
		t.Errorf("want: %+v, got: %+v", want, got)
	}
}

func TestInvalidValue(t *testing.T) {
	t.Setenv("INT", "krosch")
	conf := SomeStruct{}
	_, err := config.New(&conf)
	expected := `field Int: cannot convert "krosch" to INT`
	if err.Error() != expected {
		t.Errorf("want: %q, got: %q", expected, err.Error())
	}
}

func TestMusNew(t *testing.T) {
	defer func() { _ = recover() }()
	conf := SomeStruct{}
	config.MustNew(&conf)
}

func TestMusNewPanic(t *testing.T) {
	defer func() { _ = recover() }()
	t.Setenv("INT", "krosch")
	conf := SomeStruct{}
	config.MustNew(&conf)
	t.Error("should have caused a panic")
}
