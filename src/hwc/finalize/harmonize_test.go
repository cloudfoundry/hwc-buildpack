package finalize_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cloudfoundry/hwc-buildpack/src/hwc/finalize"

	"github.com/cloudfoundry/libbuildpack"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hamonizer", func() {
	var (
		err      error
		buildDir string
		depDir   string
		hwc      finalize.HarmonizerImpl
		buffer   *bytes.Buffer
	)

	BeforeEach(func() {
		buildDirParent, err := ioutil.TempDir("", "hwc-buildpack.build.")
		Expect(err).ToNot(HaveOccurred())

		buildDir = filepath.Join(buildDirParent, "eventual-build-dir")

		depsDir, err := ioutil.TempDir("", "hwc-buildpack.deps.")
		Expect(err).ToNot(HaveOccurred())

		depDir = filepath.Join(depsDir, "0")
		err = os.MkdirAll(depDir, 0777)
		Expect(err).To(BeNil())

		buffer = new(bytes.Buffer)
		logger := libbuildpack.NewLogger(buffer)

		hwc = finalize.HarmonizerImpl{
			BuildDir: buildDir,
			DepDir:   depDir,
			Log:      logger,
		}
	})

	Describe("CheckWebConfig", func() {
		Context("build dir does not exist", func() {
			It("returns an error", func() {
				err = hwc.CheckWebConfig()
				Expect(err.Error()).To(Equal("Invalid build directory provided"))
			})
		})

		Context("build dir exists", func() {
			BeforeEach(func() {
				err = os.MkdirAll(buildDir, 0777)
				Expect(err).To(BeNil())
			})

			Context("build dir does not contain web.config", func() {
				It("returns an error", func() {
					err = hwc.CheckWebConfig()
					Expect(err.Error()).To(Equal("Missing Web.config"))
				})
			})

			Context("build dir contains web.config", func() {
				BeforeEach(func() {
					err = ioutil.WriteFile(filepath.Join(buildDir, "Web.ConfiG"), []byte("xx"), 0644)
					Expect(err).To(BeNil())
				})

				It("does not return an error", func() {
					err = hwc.CheckWebConfig()
					Expect(err).To(BeNil())
				})
			})
		})
	})

	Describe("LinkLegacyHwc", func() {
		Context("dep dir contains hwc/hwc.exe", func() {
			BeforeEach(func() {
				hwcDepPath := filepath.Join(depDir, "hwc", "hwc.exe")
				hwc32DepPath := filepath.Join(depDir, "hwc", "hwc_x86.exe")

				err := os.MkdirAll(filepath.Dir(hwcDepPath), 0777)
				Expect(err).To(BeNil())

				err = ioutil.WriteFile(hwcDepPath, []byte("exe"), 0744)
				Expect(err).To(BeNil())

				err = ioutil.WriteFile(hwc32DepPath, []byte("exe"), 0744)
				Expect(err).To(BeNil())
			})

			It("links the file to build dir .cloudfoundry/hwc.exe", func() {
				err = hwc.LinkLegacyHwc()
				Expect(err).To(BeNil())

				Expect(filepath.Join(buildDir, ".cloudfoundry", "hwc.exe")).To(BeARegularFile())
				Expect(filepath.Join(buildDir, ".cloudfoundry", "hwc_x86.exe")).To(BeARegularFile())
			})
		})

		Context("dep dir does not contain hwc.exe", func() {
			It("returns an error", func() {
				err = hwc.LinkLegacyHwc()
				Expect(err.Error()).To(Equal("Missing hwc.exe"))
			})
		})
	})

})
