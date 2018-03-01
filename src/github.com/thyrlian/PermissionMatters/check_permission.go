package main

import (
	"./permission"
	"./storage"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
}

func getPermissions(apkFile string) []permission.Permission {
	// check $ANDROID_HOME
	androidHomePath := strings.TrimSuffix(os.Getenv("ANDROID_HOME"), "/")
	if len(androidHomePath) == 0 {
		log.Fatal("Environment variable ANDROID_HOME is not defined")
	} else {
		log.Printf("$ANDROID_HOME: %s", androidHomePath)
	}
	// check apkanalyzer
	apkanalyzerFile := androidHomePath + "/tools/bin/apkanalyzer"
	if _, err := os.Stat(apkanalyzerFile); err == nil {
		log.Printf("Found apkanalyzer at: %s", apkanalyzerFile)
	} else {
		log.Fatalf("Can not find apkanalyzer at %s", apkanalyzerFile)
	}
	// build and execute command
	args := []string{"manifest", "permissions", apkFile}
	cmd := exec.Command(apkanalyzerFile, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	permissionsRaw := strings.Split(strings.TrimSpace(string(out)), "\n")
	permissions := make([]permission.Permission, len(permissionsRaw))
	for i := range permissionsRaw {
		permissions[i] = permission.New(permissionsRaw[i])
	}
	return permissions
}
