package libError_test

import (
	"errors"
	"testing"

	libError "github.com/netsak/go-libstd/error"
)

func TestFailOnError(t *testing.T) {
	defer func() { _ = recover() }()
	err := errors.New("test-error")
	libError.FailOnError(err, "my error message")
	t.Error("should have caused a panic")
}

func TestFailOnErrorMessage(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			got := r.(error).Error()
			want := "test-error"
			if got != want {
				t.Fatalf("want: %q, got: %q", want, got)
			}
		}
	}()
	err := errors.New("test-error")
	libError.FailOnError(err, "my error message")
	t.Error("should have caused a panic")
}

func TestFailOnErrorWithoutError(t *testing.T) {
	libError.FailOnError(nil, "should not panic")
}
