package permissionguard

import (
	"io/ioutil"
)

func PersistOntoDisk(json string, file string) {
	err := ioutil.WriteFile(file, []byte(json), 0644)
	if err != nil {
		panic(err)
	}
}

func LoadFromDisk(file string) []Permission {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return Deserialize(dat)
}
