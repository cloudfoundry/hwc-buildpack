package compile_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cloudfoundry/hwc-buildpack/src/hwc/compile"

	"github.com/cloudfoundry/libbuildpack"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate mockgen -source=compile.go --destination=mocks_test.go --package=compile_test

var _ = Describe("Compile", func() {
	var (
		err           error
		buildDir      string
		compiler      compile.Compiler
		logger        *libbuildpack.Logger
		buffer        *bytes.Buffer
		mockCtrl      *gomock.Controller
		mockManifest  *MockManifest
		mockInstaller *MockInstaller
	)

	BeforeEach(func() {
		buildDir, err = ioutil.TempDir("", "hwc-buildpack.build.")
		Expect(err).To(BeNil())

		buffer = new(bytes.Buffer)
		logger = libbuildpack.NewLogger(buffer)

		mockCtrl = gomock.NewController(GinkgoT())
		mockManifest = NewMockManifest(mockCtrl)
		mockInstaller = NewMockInstaller(mockCtrl)

		compiler = compile.Compiler{
			BuildDir:  buildDir,
			Manifest:  mockManifest,
			Installer: mockInstaller,
			Log:       logger,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()

		err = os.RemoveAll(buildDir)
		Expect(err).To(BeNil())
	})

	Describe("CheckWebConfig", func() {
		Context("build dir does not exist", func() {
			BeforeEach(func() {
				compiler.BuildDir = "not/a/directory"
			})

			It("returns an error", func() {
				err = compiler.CheckWebConfig()
				Expect(err.Error()).To(Equal("Invalid build directory provided"))
			})
		})

		Context("build dir does not contain web.config", func() {
			It("returns an error", func() {
				err = compiler.CheckWebConfig()
				Expect(err.Error()).To(Equal("Missing Web.config"))
			})
		})

		Context("build dir contains web.config", func() {
			BeforeEach(func() {
				err = ioutil.WriteFile(filepath.Join(buildDir, "Web.ConfiG"), []byte("xx"), 0644)
				Expect(err).To(BeNil())
			})

			It("does not return an error", func() {
				err = compiler.CheckWebConfig()
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("InstallHWC", func() {
		It("installs HWC to <build_dir>/.cloudfoundry", func() {
			dep := libbuildpack.Dependency{Name: "hwc", Version: "99.99"}

			mockManifest.EXPECT().DefaultVersion("hwc").Return(dep, nil)
			mockInstaller.EXPECT().InstallDependency(dep, filepath.Join(buildDir, ".cloudfoundry"))

			err = compiler.InstallHWC()
			Expect(err).To(BeNil())

			Expect(buffer.String()).To(ContainSubstring("-----> Installing HWC"))
			Expect(buffer.String()).To(ContainSubstring("HWC version 99.99"))
		})
	})
})
