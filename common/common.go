package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	errInvalidBuildDir  = errors.New("Invalid build directory provided")
	errMissingWebConfig = errors.New("Missing Web.config")
)

func Fail(err error) {
	fmt.Fprintf(os.Stderr, "\n%s\n", err)
	os.Exit(1)
}

func CheckErr(err error) {
	if err != nil {
		Fail(err)
	}
}

func CheckWebConfig(buildDir string) {
	_, err := os.Stat(buildDir)
	if err != nil {
		Fail(errInvalidBuildDir)
	}

	files, err := ioutil.ReadDir(buildDir)
	if err != nil {
		Fail(errInvalidBuildDir)
	}

	var webConfigExists bool
	for _, f := range files {
		if strings.ToLower(f.Name()) == "web.config" {
			webConfigExists = true
			break
		}
	}

	if !webConfigExists {
		Fail(errMissingWebConfig)
	}
}
