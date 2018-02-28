package comparator

import (
	"../permission"
)

func getDiff(base []permission.Permission, goal []permission.Permission) []permission.Permission {
	var diff []permission.Permission
	for i := range base {
		nonexistence := true
		for j := range goal {
			if base[i] == goal[j] {
				nonexistence = false
				break
			}
		}
		if nonexistence {
			diff = append(diff, base[i])
		}
	}
	return diff
}

func GetLess(old []permission.Permission, new []permission.Permission) []permission.Permission {
	return getDiff(old, new)
}

func GetMore(old []permission.Permission, new []permission.Permission) []permission.Permission {
	return getDiff(new, old)
}
