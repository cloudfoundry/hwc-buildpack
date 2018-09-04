package finalize

import (
	"io"

	"github.com/cloudfoundry/libbuildpack"
)

type Stager interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/stager.go
	BuildDir() string
	DepDir() string
	DepsIdx() string
	DepsDir() string
	AddBinDependencyLink(string, string) error
}

type Manifest interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/manifest.go
	AllDependencyVersions(string) []string
	DefaultVersion(string) (libbuildpack.Dependency, error)
}

type Installer interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/installer.go
	InstallDependency(libbuildpack.Dependency, string) error
	InstallOnlyVersion(string, string) error
}

type Command interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/command.go
	Execute(string, io.Writer, io.Writer, string, ...string) error
	Output(dir string, program string, args ...string) (string, error)
}

type Harmonizer interface {
	CheckWebConfig() error
	LinkLegacyHwc() error
}

type Finalizer struct {
	BuildDir   string
	Manifest   Manifest
	Stager     Stager
	Command    Command
	Harmonizer Harmonizer
	Log        *libbuildpack.Logger
}

func (f *Finalizer) Run() error {
	f.Log.BeginStep("Configuring hwc")

	if err := f.Harmonizer.CheckWebConfig(); err != nil {
		f.Log.Error("Unable to locate web.config: %s", err.Error())
		return err
	}

	if err := f.Harmonizer.LinkLegacyHwc(); err != nil {
		f.Log.Error("Unable to install HWC: %s", err.Error())
		return err
	}

	return nil
}
