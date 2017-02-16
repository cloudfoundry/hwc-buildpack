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

	bpRoot, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), ".."))
	checkErr(err)

	err = Compile(os.Args[1:], bpRoot)
	checkErr(err)

	os.Exit(0)
}

func Compile(args []string, bpRoot string) error {
	buildDir, _, err := parseArgs(args)
	if err != nil {
		return err
	}

	err = checkWebConfig(buildDir)
	if err != nil {
		return err
	}

	manifest, err := bp.NewManifest(bpRoot)
	if err != nil {
		return err
	}

	defaultHWC, err := manifest.DefaultVersion("hwc")
	if err != nil {
		return err
	}

	tmpDir, err := ioutil.TempDir("", "hwc")
	if err != nil {
		return err
	}

	hwcZipFile := filepath.Join(tmpDir, "hwc.zip")

	err = manifest.FetchDependency(defaultHWC, hwcZipFile)
	if err != nil {
		return err
	}

	hwcDir := filepath.Join(buildDir, ".cloudfoundry")
	err = os.MkdirAll(hwcDir, 0700)
	if err != nil {
		return err
	}

	err = bp.ExtractZip(hwcZipFile, hwcDir)
	if err != nil {
		return err
	}

	return nil
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

func checkWebConfig(buildDir string) error {
	_, err := os.Stat(buildDir)
	if err != nil {
		return errInvalidBuildDir
	}

	files, err := ioutil.ReadDir(buildDir)
	if err != nil {
		return errInvalidBuildDir
	}

	var webConfigExists bool
	for _, f := range files {
		if strings.ToLower(f.Name()) == "web.config" {
			webConfigExists = true
			break
		}
	}

	if !webConfigExists {
		return errMissingWebConfig
	}
	return nil
}
