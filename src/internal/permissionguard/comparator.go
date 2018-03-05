package permissionguard

func getDiff(base []Permission, goal []Permission) []Permission {
	var diff []Permission
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

// GetLess compares two slices of permissions, and returns a slice of removed permissions
func GetLess(old []Permission, new []Permission) []Permission {
	return getDiff(old, new)
}

// GetMore compares two slices of permissions, and returns a slice of added permissions
func GetMore(old []Permission, new []Permission) []Permission {
	return getDiff(new, old)
}
