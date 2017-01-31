package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

const APP_PORT = "43311"
const APP_NAME = "nora"

var _ = Describe("HWC", func() {
	var (
		err        error
		binaryPath string
		tmpDir     string
	)

	BeforeEach(func() {
		binaryPath, err = Build("github.com/greenhouse-org/hwc-buildpack/hwc")
		Expect(err).ToNot(HaveOccurred())
		tmpDir, err = ioutil.TempDir("", "")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(tmpDir)
		CleanupBuildArtifacts()
	})

	sendCtrlBreak := func(s *Session) {
		d, err := syscall.LoadDLL("kernel32.dll")
		Expect(err).ToNot(HaveOccurred())
		p, err := d.FindProc("GenerateConsoleCtrlEvent")
		Expect(err).ToNot(HaveOccurred())
		r, _, err := p.Call(syscall.CTRL_BREAK_EVENT, uintptr(s.Command.Process.Pid))
		Expect(r).ToNot(Equal(0), fmt.Sprintf("GenerateConsoleCtrlEvent: %v\n", err))
	}

	startApp := func(name, port, userProfile string) (*Session, error) {
		cmd := exec.Command(binaryPath)
		cmd.Env = append(os.Environ(), fmt.Sprintf("PORT=%s", port))
		cmd.Env = append([]string{fmt.Sprintf("USERPROFILE=%s", userProfile)}, cmd.Env...)
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		cmd.Dir = filepath.Join(wd, "fixtures", APP_NAME)
		cmd.SysProcAttr = &syscall.SysProcAttr{
			CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
		}

		return Start(cmd, GinkgoWriter, GinkgoWriter)
	}

	Context("when the app PORT is not set", func() {
		It("errors", func() {
			session, err := startApp(APP_NAME, "", tmpDir)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(Exit(1))
			Eventually(session.Err).Should(Say("Missing PORT environment variable"))
		})
	})

	Context("when the app USERPROFILE is not set", func() {
		It("errors", func() {
			session, err := startApp(APP_NAME, APP_PORT, "")
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(Exit(1))
			Eventually(session.Err).Should(Say("Missing USERPROFILE environment variable"))
		})
	})

	Context("Given that I have an ASP.NET MVC application", func() {
		var (
			session *Session
			err     error
		)

		BeforeEach(func() {
			session, err = startApp("nora", APP_PORT, tmpDir)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(Say("Server Started"))
		})

		AfterEach(func() {
			sendCtrlBreak(session)
			Eventually(session, 10*time.Second).Should(Say("Server Shutdown"))
			Eventually(session).Should(Exit(0))
		})

		It("runs it on the specified port", func() {
			url := fmt.Sprintf("http://localhost:%s", APP_PORT)
			res, err := http.Get(url)
			Expect(err).ToNot(HaveOccurred())

			body, err := ioutil.ReadAll(res.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(body)).To(Equal(fmt.Sprintf(`"hello i am %s"`, APP_NAME)))
		})

		It("correctly utilizes the USERPROFILE temp directory", func() {
			url := fmt.Sprintf("http://localhost:%s", APP_PORT)
			_, err := http.Get(url)
			Expect(err).ToNot(HaveOccurred())

			_, err = os.Stat(filepath.Join(tmpDir, "tmp", "root"))
			Expect(err).ToNot(HaveOccurred())

			By("placing config files in the temp directory", func() {
				_, err = os.Stat(filepath.Join(tmpDir, "tmp", "config", "Web.config"))
				Expect(err).ToNot(HaveOccurred())
				_, err = os.Stat(filepath.Join(tmpDir, "tmp", "config", "ApplicationHost.config"))
				Expect(err).ToNot(HaveOccurred())
				_, err = os.Stat(filepath.Join(tmpDir, "tmp", "config", "Aspnet.config"))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		It("does not add unexpected custom headers", func() {
			url := fmt.Sprintf("http://localhost:%s", APP_PORT)
			res, err := http.Get(url)
			Expect(err).ToNot(HaveOccurred())

			var customHeaders []string
			for h, _ := range res.Header {
				if strings.HasPrefix(h, "X-") {
					customHeaders = append(customHeaders, strings.ToLower(h))
				}
			}
			Expect(len(customHeaders)).To(Equal(1))
			Expect(customHeaders[0]).To(Equal("x-aspnet-version"))
		})
	})

	Context("Given that I have an ASP.NET Classic application", func() {
		It("runs on the specified port", func() {
			session, err := startApp("asp-classic", APP_PORT, tmpDir)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(Say("Server Started"))

			url := fmt.Sprintf("http://localhost:%s", APP_PORT)
			res, err := http.Get(url)
			Expect(err).ToNot(HaveOccurred())

			body, err := ioutil.ReadAll(res.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(body)).To(Equal(fmt.Sprintf(`"hello i am %s"`, APP_NAME)))

			sendCtrlBreak(session)
			Eventually(session, 10*time.Second).Should(Say("Server Shutdown"))
			Eventually(session).Should(Exit(0))
		})
	})
})
