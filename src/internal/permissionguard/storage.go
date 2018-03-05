package permissionguard

import (
	"io/ioutil"
)

// PersistOntoDisk writes a JSON-encoded string to a file
func PersistOntoDisk(json string, file string) {
	err := ioutil.WriteFile(file, []byte(json), 0644)
	if err != nil {
		panic(err)
	}
}

// LoadFromDisk reads a JSON-encoded file, and returns a Permission slice
func LoadFromDisk(file string) []Permission {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return Deserialize(dat)
}
