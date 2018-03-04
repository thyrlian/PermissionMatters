package main

import (
	"../internal/permissionguard"
)

func main() {
	order := permissionguard.ParseArgs()
	if order.Operation == permissionguard.Take {
		permissions := permissionguard.GetPermissions(order.Apk)
		permissionguard.TakeSnapshot(permissions, order.Snapshot)
	} else if order.Operation == permissionguard.Scan {
		basePermissions := permissionguard.LoadFromDisk(order.Snapshot)
		goalPermissions := permissionguard.GetPermissions(order.Apk)
		conclusion := permissionguard.Check(basePermissions, goalPermissions)
		permissionguard.Process(conclusion)
	}
}
