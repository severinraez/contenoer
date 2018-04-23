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
	Bundles map[string]composefile
}

type composefile struct {
	Path string
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

	composeFile := composefile{
		Path: absPath,
		Name: name}

	result.Bundles[name] = composeFile

	return result, nil
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
		Bundles: make(map[string]composefile)}
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
