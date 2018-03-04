package permissionguard

import (
	"flag"
	"fmt"
	"os"
)

// 2 different types of subcommand
const (
	Take = iota
	Scan = iota
)

// Order tells what Operation should be performed, and with detailed information
type Order struct {
	Operation int
	Snapshot  string
	Apk       string
}

// ParseArgs parse the command line arguments and returns an Order
func ParseArgs() Order {
	var operation int
	var snapshot string
	var apk string
	takeCmd := flag.NewFlagSet("take", flag.ExitOnError)
	scanCmd := flag.NewFlagSet("scan", flag.ExitOnError)
	takeSnapshotVal := takeCmd.String("snapshot", "./permissions.json", "The permission snapshot file.")
	takeApkVal := takeCmd.String("apk", "", "The APK file to analyze.")
	scanSnapshotVal := scanCmd.String("snapshot", "./permissions.json", "The permission snapshot file.")
	scanApkVal := scanCmd.String("apk", "", "The APK file to analyze.")
	if len(os.Args) < 2 {
		fmt.Println("Require sub-command: 'take' or 'scan'")
		os.Exit(2)
	}
	switch os.Args[1] {
	case "take":
		operation = Take
		takeCmd.Parse(os.Args[2:])
	case "scan":
		operation = Scan
		scanCmd.Parse(os.Args[2:])
	default:
		fmt.Println("Supported sub-commands: 'take' or 'scan'")
		os.Exit(2)
	}
	if takeCmd.Parsed() {
		if *takeApkVal == "" {
			takeCmd.PrintDefaults()
			os.Exit(2)
		}
		snapshot = *takeSnapshotVal
		apk = *takeApkVal
	}
	if scanCmd.Parsed() {
		if *scanApkVal == "" {
			scanCmd.PrintDefaults()
			os.Exit(2)
		}
		snapshot = *scanSnapshotVal
		apk = *scanApkVal
	}
	return Order{operation, snapshot, apk}
}
