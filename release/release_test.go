package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

const expectedReleaseInfo = `---
default_process_types:
  web: .cloudfoundry\\hwc.exe`

var _ = Describe("Release", func() {
	var (
		err        error
		binaryPath string
		buildDir   string
	)

	BeforeEach(func() {
		binaryPath, err = Build("github.com/greenhouse-org/hwc-buildpack/release")
		Expect(err).ToNot(HaveOccurred())

		buildDir, err = ioutil.TempDir("", "")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(buildDir)
		CleanupBuildArtifacts()
	})

	Context("when a Web.config exists in the build directory", func() {
		It("succeeds", func() {
			err := ioutil.WriteFile(filepath.Join(buildDir, "Web.config"), []byte("XML"), 0666)
			Expect(err).ToNot(HaveOccurred())

			cmd := exec.Command(binaryPath, buildDir)
			session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(Exit(0))
			Eventually(session.Out.Contents()).Should(Equal([]byte(expectedReleaseInfo)))
		})
	})

	Context("when not provided any arguments", func() {
		It("fails", func() {
			cmd := exec.Command(binaryPath)
			session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(Exit(1))
			Eventually(session.Err).Should(Say("Invalid usage. Expected: release.exe <build_dir>"))
		})
	})

	Context("when provided a nonexistent build directory", func() {
		It("fails", func() {
			cmd := exec.Command(binaryPath, "/nonexistent/build_dir")
			session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(Exit(1))
			Eventually(session.Err).Should(Say("Invalid build directory provided"))
		})
	})

	Context("when the app does not include a Web.config", func() {
		It("fails", func() {
			cmd := exec.Command(binaryPath, buildDir)
			session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(Exit(1))
			Eventually(session.Err).Should(Say("Missing Web.config"))
		})
	})
})
