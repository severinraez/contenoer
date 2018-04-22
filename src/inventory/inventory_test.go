package inventory

import (
	"testing"
	"reflect"
//	"fmt"
	"os"
)

func TestAdd(t *testing.T) {
	i := New()

	path := "/tmp/composefile"
	os.Create(path)

	i, err := Add(i, "name", path)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	want := []string{"name"}
	got := BundleNames(i)

	if ! reflect.DeepEqual(want, got) {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func TestNew(t *testing.T) {
	i := New()

	bundles := BundleNames(i)

	if len(bundles) != 0 {
		t.Errorf("Bundles not empty: %v", bundles)
	}
}
