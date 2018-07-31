package integration_test

import (
	"fmt"
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
