package inventory

import (
	"fmt"
	"encoding/json"
	"github.com/jinzhu/copier"
	"path/filepath"
	"errors"
	"os"
)

type Inventory struct {
	Bundles map[string]Bundle
}

type Bundle struct {
	ComposeFilePath string
	Name string
}


func Add(i Inventory, name string, path string) (Inventory, error){
	if ! isDockerfile(path) {
		err := errors.New(fmt.Sprintf("No Dockerfile at %s", path))
		return Inventory{}, err
	}

	result := Inventory{}
	copier.Copy(&result, &i)

	absPath, _ := filepath.Abs(path)

	bundle := Bundle{
		ComposeFilePath: absPath,
		Name: name}

	result.Bundles[name] = bundle

	return result, nil
}

func Bundles(i Inventory) []Bundle {
	var bundles []Bundle
	for _, bundle := range i.Bundles {
		bundles = append(bundles, bundle)
	}
	return bundles
}

func BundleNames(i Inventory) []string {
	var names []string
	for _, bundle := range i.Bundles {
		names = append(names, bundle.Name)
	}
	return names
}

func New() Inventory {
	return Inventory{
		Bundles: make(map[string]Bundle)}
}

func Serialize(i Inventory) ([]byte, error) {
	json, err := json.Marshal(i)
	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func Deserialize(jsonBlob []byte) (Inventory, error) {
	i := Inventory{}
	err := json.Unmarshal(jsonBlob, &i)

	return i, err
}

func isDockerfile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
