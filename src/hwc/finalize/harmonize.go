package finalize

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudfoundry/libbuildpack"
)

// move to hwc package and rename to just Hwc
type HarmonizerImpl struct {
	Log      *libbuildpack.Logger
	BuildDir string
	DepDir   string
}

var (
	errInvalidBuildDir  = errors.New("Invalid build directory provided")
	errMissingWebConfig = errors.New("Missing Web.config")
	errMissingDepHwc    = errors.New("Missing hwc.exe")
)

func NewHarmonizer(logger *libbuildpack.Logger, buildDir, depDir string) *HarmonizerImpl {
	return &HarmonizerImpl{Log: logger, BuildDir: buildDir, DepDir: depDir}
}

func (h *HarmonizerImpl) CheckWebConfig() error {
	_, err := os.Stat(h.BuildDir)
	if err != nil {
		return errInvalidBuildDir
	}

	files, err := ioutil.ReadDir(h.BuildDir)
	if err != nil {
		return errInvalidBuildDir
	}

	var webConfigExists bool
	for _, file := range files {
		if strings.ToLower(file.Name()) == "web.config" {
			webConfigExists = true
			break
		}
	}

	if !webConfigExists {
		return errMissingWebConfig
	}
	return nil
}

func (h *HarmonizerImpl) LinkLegacyHwc() error {
	sourceHwcPath := filepath.Join(h.DepDir, "hwc", "hwc.exe")
	if _, err := os.Stat(sourceHwcPath); err != nil {
		return errMissingDepHwc
	}

	legacyHwcDir := filepath.Join(h.BuildDir, ".cloudfoundry")
	if err := os.MkdirAll(legacyHwcDir, 0777); err != nil {
		return err
	}

	legacyHwcPath := filepath.Join(legacyHwcDir, "hwc.exe")

	if err := os.Link(sourceHwcPath, legacyHwcPath); err != nil {
		h.Log.Error("Unable to install HWC: %s", err.Error())
		return err
	}

	return nil
}
