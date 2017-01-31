package main

import (
	"errors"
	"fmt"
	"os"

	. "github.com/greenhouse-org/hwc-buildpack/common"
)

const releaseInfo = `---
default_process_types:
  web: .cloudfoundry\\hwc.exe`

func main() {
	buildDir, err := parseArgs(os.Args[1:])
	CheckErr(err)

	CheckWebConfig(buildDir)

	fmt.Printf(releaseInfo)

	os.Exit(0)
}

func parseArgs(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("Invalid usage. Expected: release.exe <build_dir>")
	}

	return args[0], nil
}
