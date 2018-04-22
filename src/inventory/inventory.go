package inventory

import (
	"fmt"
	"json"
)

type inventory struct {
	bundles composefile[]
}

type composefile struct {
	path string
}

func add(i inventory, path string) {
	if ! isDockerfile(path) {
		err = fmt.Sprintf("No Dockerfile at %s", path)
		return [nil, err]
	}

	return inventory{
		i.bundles += composefile{
			path: path
		}
	}
}

func initialize() {
	return new(inventory)
}

func serialize(i inventory) byte[] {
	json, err := json.Marshal(i)
	return json
}

func deserialize(json string) inventory {
	i = new(inventory)
	err := json.Unmarshal(json, &i)

	return i
}

func isDockerfile(path) {
	_, err := os.Stat("/path/to/whatever")
	return err == nil
}
