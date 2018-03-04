package analyzer

import (
	"../permission"
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func GetPermissions(apkFile string) []permission.Permission {
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
	// check the existence of the given APK file
	if _, err := os.Stat(apkFile); os.IsNotExist(err) {
		log.Fatalf("The given APK file '%s' doesn't exist", apkFile)
	}
	// build and execute command
	args := []string{"manifest", "permissions", apkFile}
	cmd := exec.Command(apkanalyzerFile, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(string(stderr.Bytes()))
	}
	permissionsRaw := strings.Split(strings.TrimSpace(string(stdout.Bytes())), "\n")
	permissions := make([]permission.Permission, len(permissionsRaw))
	for i := range permissionsRaw {
		permissions[i] = permission.New(permissionsRaw[i])
	}
	return permissions
}
