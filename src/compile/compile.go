package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	bp "github.com/cloudfoundry/libbuildpack"
)

func main() {
	buildDir, _, err := parseArgs(os.Args[1:])
	checkErr(err)

	checkWebConfig(buildDir)

	bpRoot, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), ".."))
	checkErr(err)

	manifest, err := bp.NewManifest(filepath.Join(bpRoot, "manifest.yml"))
	checkErr(err)

	defaultHWC, err := manifest.DefaultVersion("hwc")
	checkErr(err)

	err = manifest.FetchDependency(defaultHWC, "/tmp/hwc.zip")
	checkErr(err)

	hwcDir := filepath.Join(buildDir, ".cloudfoundry")
	err = os.MkdirAll(hwcDir, 0700)
	checkErr(err)

	err = bp.ExtractZip("/tmp/hwc.zip", hwcDir)
	checkErr(err)

	os.Exit(0)
}

func parseArgs(args []string) (string, string, error) {
	if len(args) != 2 {
		return "", "", errors.New("Invalid usage. Expected: compile.exe <build_dir> <cache_dir>")
	}

	return args[0], args[1], nil
}

var (
	errInvalidBuildDir  = errors.New("Invalid build directory provided")
	errMissingWebConfig = errors.New("Missing Web.config")
)

func fail(err error) {
	fmt.Fprintf(os.Stderr, "\n%s\n", err)
	os.Exit(1)
}

func checkErr(err error) {
	if err != nil {
		fail(err)
	}
}

func checkWebConfig(buildDir string) {
	_, err := os.Stat(buildDir)
	if err != nil {
		fail(errInvalidBuildDir)
	}

	files, err := ioutil.ReadDir(buildDir)
	if err != nil {
		fail(errInvalidBuildDir)
	}

	var webConfigExists bool
	for _, f := range files {
		if strings.ToLower(f.Name()) == "web.config" {
			webConfigExists = true
			break
		}
	}

	if !webConfigExists {
		fail(errMissingWebConfig)
	}
}
