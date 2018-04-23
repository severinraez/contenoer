package inventory

import (
	"testing"
	"os"
	. "github.com/severinraez/cotenoer/testhelper"
)

func TestNew(t *testing.T) {
	i := New()

	bundles := BundleNames(i)

	AssertEqual(len(bundles), 0, t)
}

func TestAdd(t *testing.T) {
	i := New()

	path := "/tmp/composefile"
	os.Create(path)

	i, err := Add(i, "name", path)

	AssertEqual(err, nil, t)

	want := []string{"name"}
	got := BundleNames(i)

	AssertEqual(got, want, t)
}

func fakeInventory() Inventory {
	path := "/tmp/composefile"
	os.Create(path)

	i, _ := Add(New(), "name", path)
	return i
}

func TestSerialize(t *testing.T) {
	i := fakeInventory()

	got, err := Serialize(i)
	want := []byte("{\"Bundles\":{\"name\":{\"Path\":\"/tmp/composefile\",\"Name\":\"name\"}}}")

	AssertEqual(err, nil, t)
	AssertEqual(string(got), string(want), t)
}

func TestDeserialize(t *testing.T) {
	serialized := []byte("{\"Bundles\":{\"name\":{\"Path\":\"/tmp/composefile\",\"Name\":\"name\"}}}")

	got, err := Deserialize(serialized)
	want := fakeInventory()

	AssertEqual(err, nil, t)
	AssertEqual(got, want, t)
}
