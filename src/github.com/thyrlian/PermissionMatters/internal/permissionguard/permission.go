package permissionguard

import (
	"encoding/json"
	"log"
	"strings"
)

// Permission represents the Android application's permission, with a simple string Name property
type Permission struct {
	Name string
}

// NewPermission instantiates a Permission by a string
func NewPermission(name string) Permission {
	permission := Permission{strings.TrimSpace(name)}
	return permission
}

func serializeToJson(object interface{}) string {
	b, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

// ToJson parses a Permission and returns a JSON-encoded string
func (p Permission) ToJson() string {
	return serializeToJson(p)
}

// ToJsonFromList parses a Permission slice and returns a JSON-encoded string
func ToJsonFromList(permissions []Permission) string {
	return serializeToJson(permissions)
}

// Deserialize parses a JSON-encoded byte slice and returns a Permission slice
func Deserialize(byt []byte) []Permission {
	var dat []Permission
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	return dat
}
