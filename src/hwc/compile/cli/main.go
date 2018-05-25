package main

import (
	"hwc/compile"
	"os"
	"time"

	"github.com/cloudfoundry/libbuildpack"
)

func main() {
	logger := libbuildpack.NewLogger(os.Stdout)

	buildpackDir, err := libbuildpack.GetBuildpackDir()
	if err != nil {
		logger.Error("Unable to determine buildpack directory: %s", err.Error())
		os.Exit(9)
	}

	manifest, err := libbuildpack.NewManifest(buildpackDir, logger, time.Now())
	if err != nil {
		logger.Error("Unable to load buildpack manifest: %s", err.Error())
		os.Exit(10)
	}
	installer := libbuildpack.NewInstaller(manifest)

	stager := libbuildpack.NewStager(os.Args[1:], logger, manifest)
	err = stager.CheckBuildpackValid()
	if err != nil {
		panic(err)
	}

	hc := compile.Compiler{
		BuildDir:  stager.BuildDir(),
		Manifest:  manifest,
		Installer: installer,
		Log:       logger,
	}

	err = hc.Compile()
	if err != nil {
		panic(err)
	}

	stager.StagingComplete()
}
