package integration_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CF HWC Buildpack", func() {
	var (
		app        *cutlass.App
		checkStack func(string)
	)

	BeforeEach(func() {
		checkStack = func(s string) {
			if !(s == os.Getenv("CF_STACK") || os.Getenv("CF_STACK") == "") {
				Skip(fmt.Sprintf("Only runs against %s or all stacks", s))
			}
		}
	})

	AfterEach(func() { app = DestroyApp(app) })

	Describe("deploying an hwc app", func() {
		Context("windows2012R2", func() {
			BeforeEach(func() {
				checkStack("windows2012R2")
				app = cutlass.New(filepath.Join(bpDir, "fixtures", "windows_app"))
				app.Stack = "windows2012R2"
			})

			It("deploys successfully", func() {
				PushAppAndConfirm(app)
				if cutlass.Cached {
					Expect(app.Stdout.String()).ToNot(ContainSubstring("Download ["))
					Expect(app.Stdout.String()).To(ContainSubstring("Copy ["))
				} else {
					Expect(app.Stdout.String()).To(ContainSubstring("Download ["))
					Expect(app.Stdout.String()).ToNot(ContainSubstring("Copy ["))
				}

				Expect(app.GetBody("/")).To(ContainSubstring("hello i am nora"))
			})
		})

		Context("windows2016", func() {
			BeforeEach(func() {
				checkStack("windows2016")
				app = cutlass.New(filepath.Join(bpDir, "fixtures", "windows_app_with_rewrite"))
				app.Stack = "windows2016"
			})

			It("deploys successfully with a rewrite rule", func() {
				PushAppAndConfirm(app)
				if cutlass.Cached {
					Expect(app.Stdout.String()).ToNot(ContainSubstring("Download ["))
					Expect(app.Stdout.String()).To(ContainSubstring("Copy ["))
				} else {
					Expect(app.Stdout.String()).To(ContainSubstring("Download ["))
					Expect(app.Stdout.String()).ToNot(ContainSubstring("Copy ["))
				}
				Expect(app.GetBody("/rewrite")).To(ContainSubstring("hello i am nora"))
			})
		})

		Context("http compression", func() {
			BeforeEach(func() {
				app = cutlass.New(filepath.Join(bpDir, "fixtures", "windows_app"))
				app.Stack = os.Getenv("CF_STACK")
			})
			It("gzip the response", func() {
				PushAppAndConfirm(app)
				//start the static cache
				app.GetBody("/Content/Site.css")

				url, err := app.GetUrl("/Content/Site.css")
				Expect(err).NotTo(HaveOccurred())

				headersUncompressed, err := GetResponseHeaders(url, map[string]string{})
				Expect(err).NotTo(HaveOccurred())

				Expect(headersUncompressed["Content-Encoding"]).NotTo(Equal([]string{"gzip"}))
				Expect(headersUncompressed["Content-Length"]).NotTo(BeEmpty())
				uncompressedLength, err := strconv.Atoi(headersUncompressed["Content-Length"][0])
				Expect(err).NotTo(HaveOccurred())

				headersCompressed, err := GetResponseHeaders(url, map[string]string{"Accept-Encoding": "gzip"})
				Expect(err).NotTo(HaveOccurred())

				Expect(headersCompressed["Content-Encoding"]).To(Equal([]string{"gzip"}))
				Expect(headersCompressed["Content-Length"]).NotTo(BeEmpty())
				compressedLength, err := strconv.Atoi(headersCompressed["Content-Length"][0])
				Expect(err).NotTo(HaveOccurred())

				Expect(compressedLength).To(BeNumerically("<", uncompressedLength))
			})
		})
	})

	// To run this test add the cryptoki.dll and Pkcs11Interop.dll to fixtures/luna_windows_app/Bin/
	// and add a ups.json file to fixtures/luna_windows_app/ with the Luna User Provided Service JSON
	// credentials/configuration.
	Describe("deploying a Luna hwc app", func() {

		var (
			upsJsonFile string
			appPath     string
		)

		Context("windows2012R2", func() {
			BeforeEach(func() {
				checkStack("windows2012R2")
				appPath = filepath.Join(bpDir, "fixtures", "luna_windows_app")
				app = cutlass.New(appPath)
				app.Stack = "windows2012R2"
				app.Buildpacks = []string{"hwc_buildpack"}
				checkFileExists := func(filePath string) {
					if _, err := os.Stat(filePath); err != nil {
						if os.IsNotExist(err) {
							Skip(fmt.Sprintf("To run the Luna test you need to add: %s", filePath))
						}
					}
				}
				upsJsonFile = filepath.Join(appPath, "ups.json")
				checkFileExists(filepath.Join(appPath, "Bin", "cryptoki.dll"))
				checkFileExists(filepath.Join(appPath, "Bin", "Pkcs11Interop.dll"))
				checkFileExists(upsJsonFile)
			})

			AfterEach(func() {
				app = DestroyApp(app)
				DeleteService("hwc-luna-ups")
			})

			It("deploys successfully", func() {
				credentials, _ := ioutil.ReadFile(upsJsonFile)
				Expect(CreateUserProvidedService("hwc-luna-ups", string(credentials))).To(Succeed())
				PushAppAndConfirm(app)
				encrypted, _ := app.GetBody("/?command=encrypt&iv=00112233445566778899AABBCCDDEEFF&data=encrypted_with_luna_hsm")
				Expect(app.GetBody("/?command=decrypt&iv=00112233445566778899AABBCCDDEEFF&data=" + encrypted)).To(ContainSubstring("encrypted_with_luna_hsm"))
			})
		})

	})
})

func GetResponseHeaders(url string, headers map[string]string) (map[string][]string, error) {
	tr := &http.Transport{
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Header, nil
}
