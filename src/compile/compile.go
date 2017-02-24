package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	bp "github.com/cloudfoundry/libbuildpack"
)

type HWCCompiler struct {
	Compiler *bp.Compiler
}

func main() {
	buildDir := os.Args[1]
	cacheDir := os.Args[2]

	logger := bp.NewLogger()

	compiler, err := bp.NewCompiler(buildDir, cacheDir, logger)
	err = compiler.CheckBuildpackValid()
	if err != nil {
		panic(err)
	}

	hc := HWCCompiler{
		Compiler: compiler,
	}

	err = hc.Compile()
	if err != nil {
		panic(err)
	}

	compiler.StagingComplete()
}

func (c *HWCCompiler) Compile() error {
	err := c.CheckWebConfig()
	if err != nil {
		c.Compiler.Log.Error("Unable to locate web.config: %s", err.Error())
		return err
	}

	err = c.InstallHWC()
	if err != nil {
		c.Compiler.Log.Error("Unable to install HWC: %s", err.Error())
		return err
	}

	return nil
}

var (
	errInvalidBuildDir  = errors.New("Invalid build directory provided")
	errMissingWebConfig = errors.New("Missing Web.config")
)

func (c *HWCCompiler) CheckWebConfig() error {
	_, err := os.Stat(c.Compiler.BuildDir)
	if err != nil {
		return errInvalidBuildDir
	}

	files, err := ioutil.ReadDir(c.Compiler.BuildDir)
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

func (c *HWCCompiler) InstallHWC() error {
	c.Compiler.Log.BeginStep("Installing HWC")

	defaultHWC, err := c.Compiler.Manifest.DefaultVersion("hwc")
	if err != nil {
		return err
	}

	c.Compiler.Log.Info("HWC version %s", defaultHWC.Version)

	hwcDir := filepath.Join(c.Compiler.BuildDir, ".cloudfoundry")
	err = os.MkdirAll(hwcDir, 0700)
	if err != nil {
		return err
	}

	return c.Compiler.Manifest.InstallDependency(defaultHWC, hwcDir)
}
