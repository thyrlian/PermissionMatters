package storage

import (
	"../permission"
	"io/ioutil"
)

func PersistOntoDisk(json string, file string) {
	err := ioutil.WriteFile(file, []byte(json), 0644)
	if err != nil {
		panic(err)
	}
}

func LoadFromDisk(file string) []permission.Permission {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return permission.Deserialize(dat)
}
