package main_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	compile "compile"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Compile", func() {
	var (
		err      error
		buildDir string
		bpRoot   string
		cwd      string
		args     []string
	)

	BeforeEach(func() {
		cwd, err = os.Getwd()
		bpRoot = filepath.Join(cwd, "fixtures", "spec_manifest")

		buildDir, err = ioutil.TempDir("", "")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(buildDir)
		CleanupBuildArtifacts()
	})

	It("places the web app server in <build_dir>/.cloudfoundry", func() {

		err := ioutil.WriteFile(filepath.Join(buildDir, "Web.config"), []byte("XML"), 0666)
		Expect(err).ToNot(HaveOccurred())

		hwcDestPath := filepath.Join(buildDir, ".cloudfoundry", "hwc.exe")

		args = []string{buildDir, "cache_dir"}
		compile.Compile(args, bpRoot)

		_, err = os.Stat(hwcDestPath)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("when not provided any arguments", func() {
		It("fails", func() {
			args = []string{}
			err = compile.Compile(args, bpRoot)
			Expect(err).ToNot(BeNil())

			Expect(err.Error()).To(Equal("Invalid usage. Expected: compile.exe <build_dir> <cache_dir>"))
		})
	})

	Context("when provided a nonexistent build directory", func() {
		It("fails", func() {
			args = []string{"/nonexistent/build_dir", "/cache_dir"}
			err = compile.Compile(args, bpRoot)
			Expect(err).ToNot(BeNil())

			Expect(err.Error()).To(Equal("Invalid build directory provided"))
		})
	})

	Context("when the app does not include a Web.config", func() {
		It("fails", func() {
			args = []string{buildDir, "/cache_dir"}
			err = compile.Compile(args, bpRoot)
			Expect(err).ToNot(BeNil())

			Expect(err.Error()).To(Equal("Missing Web.config"))
		})
	})
})
