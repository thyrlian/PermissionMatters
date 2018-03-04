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

func GetLess(old []Permission, new []Permission) []Permission {
	return getDiff(old, new)
}

func GetMore(old []Permission, new []Permission) []Permission {
	return getDiff(new, old)
}
