package main

import (
	"errors"
	"os"
	"path/filepath"

	. "github.com/greenhouse-org/hwc-buildpack/common"
)

func main() {
	buildDir, _, err := parseArgs(os.Args[1:])
	CheckErr(err)

	CheckWebConfig(buildDir)

	hwcDir := filepath.Join(buildDir, ".cloudfoundry")
	err = os.MkdirAll(hwcDir, 0700)
	CheckErr(err)

	binDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	CheckErr(err)

	err = os.Rename(filepath.Join(binDir, "hwc.exe"), filepath.Join(hwcDir, "hwc.exe"))
	CheckErr(err)

	os.Exit(0)
}

func parseArgs(args []string) (string, string, error) {
	if len(args) != 2 {
		return "", "", errors.New("Invalid usage. Expected: compile.exe <build_dir> <cache_dir>")
	}

	return args[0], args[1], nil
}
