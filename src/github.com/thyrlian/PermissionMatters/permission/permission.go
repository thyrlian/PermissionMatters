package permission

import (
	"encoding/json"
	"log"
	"strings"
)

type Permission struct {
	Name string
}

func New(name string) Permission {
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
