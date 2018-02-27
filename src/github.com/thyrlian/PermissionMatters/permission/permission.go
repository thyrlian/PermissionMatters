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

func (p Permission) ToJson() string {
	b, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func SerializeListToJson(permissions []Permission) string {
	b, err := json.Marshal(permissions)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
