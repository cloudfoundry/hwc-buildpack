package main_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	compile "compile"

	bp "github.com/cloudfoundry/libbuildpack"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate mockgen -source=vendor/github.com/cloudfoundry/libbuildpack/manifest.go --destination=mocks_manifest_test.go --package=main_test --imports=.=github.com/cloudfoundry/libbuildpack

var _ = Describe("Compile", func() {
	var (
		err          error
		buildDir     string
		cacheDir     string
		stager       compile.HWCCompiler
		logger       bp.Logger
		buffer       *bytes.Buffer
		mockCtrl     *gomock.Controller
		mockManifest *MockManifest
	)

	BeforeEach(func() {
		buildDir, err = ioutil.TempDir("", "hwc-buildpack.build.")
		Expect(err).To(BeNil())

		cacheDir, err = ioutil.TempDir("", "hwc-buildpack.cache.")
		Expect(err).To(BeNil())
		buffer = new(bytes.Buffer)

		logger = bp.NewLogger()
		logger.SetOutput(buffer)

		mockCtrl = gomock.NewController(GinkgoT())
		mockManifest = NewMockManifest(mockCtrl)
	})

	AfterEach(func() {
		err = os.RemoveAll(buildDir)
		Expect(err).To(BeNil())

		err = os.RemoveAll(cacheDir)
		Expect(err).To(BeNil())
	})

	JustBeforeEach(func() {
		bpc := bp.Stager{
			BuildDir: buildDir,
			CacheDir: cacheDir,
			Manifest: mockManifest,
			Log:      logger,
		}

		stager = compile.HWCCompiler{Stager: &bpc}
	})

	Describe("InstallHWC", func() {})

	Describe("CheckWebConfig", func() {
		Context("build dir does not exist", func() {
			BeforeEach(func() {
				buildDir = "not/a/directory"
			})

			It("returns an error", func() {
				err = stager.CheckWebConfig()
				Expect(err.Error()).To(Equal("Invalid build directory provided"))
			})
		})

		Context("build dir does not contain web.config", func() {
			It("returns an error", func() {
				err = stager.CheckWebConfig()
				Expect(err.Error()).To(Equal("Missing Web.config"))
			})
		})

		Context("build dir contains web.config", func() {
			BeforeEach(func() {
				err = ioutil.WriteFile(filepath.Join(buildDir, "Web.ConfiG"), []byte("xx"), 0644)
				Expect(err).To(BeNil())
			})

			It("does not return an error", func() {
				err = stager.CheckWebConfig()
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("InstallHWC", func() {
		It("installs HWC to <build_dir>/.cloudfoundry", func() {
			dep := bp.Dependency{Name: "hwc", Version: "99.99"}

			mockManifest.EXPECT().DefaultVersion("hwc").Return(dep, nil)
			mockManifest.EXPECT().InstallDependency(dep, filepath.Join(buildDir, ".cloudfoundry"))

			err = stager.InstallHWC()
			Expect(err).To(BeNil())

			Expect(buffer.String()).To(ContainSubstring("-----> Installing HWC"))
			Expect(buffer.String()).To(ContainSubstring("HWC version 99.99"))
		})
	})
})
