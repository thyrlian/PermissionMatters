package permissionguard

import (
	"encoding/json"
	"log"
	"strings"
)

type Permission struct {
	Name string
}

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

func (p Permission) ToJson() string {
	return serializeToJson(p)
}

func ToJsonFromList(permissions []Permission) string {
	return serializeToJson(permissions)
}

func Deserialize(byt []byte) []Permission {
	var dat []Permission
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	return dat
}
