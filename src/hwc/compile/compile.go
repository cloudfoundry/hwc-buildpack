package compile

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudfoundry/libbuildpack"
)

type Manifest interface {
	DefaultVersion(string) (libbuildpack.Dependency, error)
}

type Installer interface {
	InstallDependency(libbuildpack.Dependency, string) error
}

type Compiler struct {
	BuildDir  string
	Manifest  Manifest
	Installer Installer
	Log       *libbuildpack.Logger
}

func (c *Compiler) Compile() error {
	err := c.CheckWebConfig()
	if err != nil {
		c.Log.Error("Unable to locate web.config: %s", err.Error())
		return err
	}

	err = c.InstallHWC()
	if err != nil {
		c.Log.Error("Unable to install HWC: %s", err.Error())
		return err
	}

	return nil
}

var (
	errInvalidBuildDir  = errors.New("Invalid build directory provided")
	errMissingWebConfig = errors.New("Missing Web.config")
)

func (c *Compiler) CheckWebConfig() error {
	_, err := os.Stat(c.BuildDir)
	if err != nil {
		return errInvalidBuildDir
	}

	files, err := ioutil.ReadDir(c.BuildDir)
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

func (c *Compiler) InstallHWC() error {
	c.Log.BeginStep("Installing HWC")

	defaultHWC, err := c.Manifest.DefaultVersion("hwc")
	if err != nil {
		return err
	}

	c.Log.Info("HWC version %s", defaultHWC.Version)

	hwcDir := filepath.Join(c.BuildDir, ".cloudfoundry")

	return c.Installer.InstallDependency(defaultHWC, hwcDir)
}
