package inventory

import (
	"fmt"
	"encoding/json"
	"github.com/jinzhu/copier"
	"path/filepath"
	"errors"
	"os"
)

type inventory struct {
	Bundles []composefile
}

type composefile struct {
	Path string
	Name string
}


func Add(i inventory, name string, path string) (inventory, error){
	if ! isDockerfile(path) {
		err := errors.New(fmt.Sprintf("No Dockerfile at %s", path))
		return inventory{}, err
	}

	result := inventory{}
	copier.Copy(&result, &i)

	absPath, _ := filepath.Abs(path)

	composeFile := composefile{
		Path: absPath,
		Name: name}

	result.Bundles = append(
		result.Bundles, composeFile)

	return result, nil
}

func BundleNames(i inventory) []string {
	var names []string
	for _, bundle := range i.Bundles {
		names = append(names, bundle.Name)
	}
	return names
}

func New() inventory {
	return inventory{}
}

func Serialize(i inventory) ([]byte, error) {
	json, err := json.Marshal(i)
	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func Deserialize(jsonBlob []byte) (inventory, error) {
	i := inventory{}
	err := json.Unmarshal(jsonBlob, &i)

	return i, err
}

func isDockerfile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
