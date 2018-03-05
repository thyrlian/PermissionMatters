package permissionguard

// 3 different types of result
const (
	Pass = iota
	Fail = iota
	Warn = iota
)

// Conclusion represents the result of a comparison, and the details of the difference
// Less tells which permissions are newly removed (which is usually good)
// More tells which permissions are newly added (which is usually bad)
type Conclusion struct {
	Result int
	Less   []Permission
	More   []Permission
}

// Check compares two slices of permissions, and return a Conclusion
func Check(base []Permission, goal []Permission) Conclusion {
	less := GetLess(base, goal)
	more := GetMore(base, goal)
	if len(more) > 0 {
		return Conclusion{Fail, less, more}
	}
	if len(less) > 0 {
		return Conclusion{Warn, less, more}
	}
	return Conclusion{Pass, less, more}
}
