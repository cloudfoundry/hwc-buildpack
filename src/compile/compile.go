package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	bp "github.com/cloudfoundry/libbuildpack"
)

type HWCStager struct {
	Stager *bp.Stager
}

func main() {
	logger := bp.NewLogger()

	stager, err := bp.NewStager(os.Args[1:], logger)
	err = stager.CheckBuildpackValid()
	if err != nil {
		panic(err)
	}

	hc := HWCStager{
		Stager: stager,
	}

	err = hc.Stage()
	if err != nil {
		panic(err)
	}

	stager.StagingComplete()
}

func (c *HWCStager) Stage() error {
	err := c.CheckWebConfig()
	if err != nil {
		c.Stager.Log.Error("Unable to locate web.config: %s", err.Error())
		return err
	}

	err = c.InstallHWC()
	if err != nil {
		c.Stager.Log.Error("Unable to install HWC: %s", err.Error())
		return err
	}

	return nil
}

var (
	errInvalidBuildDir  = errors.New("Invalid build directory provided")
	errMissingWebConfig = errors.New("Missing Web.config")
)

func (c *HWCStager) CheckWebConfig() error {
	_, err := os.Stat(c.Stager.BuildDir)
	if err != nil {
		return errInvalidBuildDir
	}

	files, err := ioutil.ReadDir(c.Stager.BuildDir)
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

func (c *HWCStager) InstallHWC() error {
	c.Stager.Log.BeginStep("Installing HWC")

	defaultHWC, err := c.Stager.Manifest.DefaultVersion("hwc")
	if err != nil {
		return err
	}

	c.Stager.Log.Info("HWC version %s", defaultHWC.Version)

	hwcDir := filepath.Join(c.Stager.BuildDir, ".cloudfoundry")

	return c.Stager.Manifest.InstallDependency(defaultHWC, hwcDir)
}
