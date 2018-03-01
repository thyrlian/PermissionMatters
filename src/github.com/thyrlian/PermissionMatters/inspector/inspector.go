package inspector

import (
	"../comparator"
	"../permission"
)

const (
	Pass              = iota
	Fail              = iota
	PassWithAttention = iota
)

type Conclusion struct {
	Result int
	Less   []permission.Permission
	More   []permission.Permission
}

func Check(base []permission.Permission, goal []permission.Permission) Conclusion {
	less := comparator.GetLess(base, goal)
	more := comparator.GetMore(base, goal)
	if len(more) > 0 {
		return Conclusion{Fail, less, more}
	}
	if len(less) > 0 {
		return Conclusion{PassWithAttention, less, more}
	}
	return Conclusion{Pass, less, more}
}
