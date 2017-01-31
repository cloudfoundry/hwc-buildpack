package main

import (
	"errors"
	"os"

	. "github.com/greenhouse-org/hwc-buildpack/common"
)

func main() {
	buildDir, err := parseArgs(os.Args[1:])
	CheckErr(err)

	CheckWebConfig(buildDir)

	os.Exit(0)
}

func parseArgs(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("Invalid usage. Expected: detect.exe <build_dir>")
	}

	return args[0], nil
}
