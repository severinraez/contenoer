package testhelper

import (
	"testing"
	"reflect"
)

func AssertEqual(got interface{}, want interface{}, t *testing.T) {
	if ! reflect.DeepEqual(got, want) {
		t.Errorf("Want %v, got %v", want, got)
	}
}
