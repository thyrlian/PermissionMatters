package main

import (
	"./analyzer"
	"./commander"
	"./handler"
	"./inspector"
	"./storage"
)

func main() {
	order := commander.ParseArgs()
	if order.Operation == commander.Take {
		permissions := analyzer.GetPermissions(order.Apk)
		handler.TakeSnapshot(permissions, order.Snapshot)
	} else if order.Operation == commander.Scan {
		basePermissions := storage.LoadFromDisk(order.Snapshot)
		goalPermissions := analyzer.GetPermissions(order.Apk)
		conclusion := inspector.Check(basePermissions, goalPermissions)
		handler.Process(conclusion)
	}
}
