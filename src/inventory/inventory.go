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
	bundles []composefile
}

type composefile struct {
	path string
	name string
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
		path: absPath,
		name: name}

	result.bundles = append(
		result.bundles, composeFile)

	return result, nil
}

func BundleNames(i inventory) []string {
	var names []string
	for _, bundle := range i.bundles {
		names = append(names, bundle.name)
	}
	return names
}

func New() inventory {
	return inventory{}
}

func Serialize(i inventory) []byte {
	json, _ := json.Marshal(i)

	return json
}

func Deserialize(jsonBlob []byte) inventory {
	i := inventory{}
	json.Unmarshal(jsonBlob, &i)

	return i
}

func isDockerfile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
