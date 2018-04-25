package testhelper

import (
	"testing"
	"reflect"
	"runtime/debug"
)

func AssertEqual(got interface{}, want interface{}, t *testing.T) {
	if ! reflect.DeepEqual(got, want) {
		t.Errorf("Want %v, got %v\n\n%s\n", want, got, debug.Stack())
	}
}

func AssertEqualH(got interface{}, want interface{}, hint string, t *testing.T) {
	if ! reflect.DeepEqual(got, want) {
		t.Errorf("%s: Want %v, got %v", hint, want, got)
	}
}
