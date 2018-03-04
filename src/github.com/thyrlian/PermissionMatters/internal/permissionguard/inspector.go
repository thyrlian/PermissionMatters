package permissionguard

const (
	Pass = iota
	Fail = iota
	Warn = iota
)

type Conclusion struct {
	Result int
	Less   []Permission
	More   []Permission
}

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
