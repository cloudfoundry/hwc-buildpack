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

		Context("with an extension buildpack", func() {
			var (
				extensionBuildpack                         cutlass.VersionedBuildpackPackage
				extensionBuildpackName, finalBuildpackName string
				err                                        error
			)

			BeforeEach(func() {
				if !ApiHasCAPIFixForMultiBP() {
					Skip("API does not have CAPI Fix for multi-buidpack. see #160325184")
				}

				if !ApiHasMultiBuildpack() {
					Skip("API does not have multi buildpack support")
				}

				// re-use the suite-built hwc_buildpack as an extension buildpack
				// this prevents having to fixture an entire buildpack to test
				finalBuildpackName = "final" + cutlass.RandStringRunes(20)
				extensionBuildpackName = "extension" + cutlass.RandStringRunes(20)

				extensionBuildpack, err = cutlass.PackageUniquelyVersionedBuildpackExtra(extensionBuildpackName, packagedBuildpack.Version, "", cutlass.Cached, ApiHasStackAssociation())
				Expect(err).NotTo(HaveOccurred())

				err = cutlass.CreateOrUpdateBuildpack(finalBuildpackName, extensionBuildpack.File, "")
				Expect(err).NotTo(HaveOccurred())

				app = cutlass.New(filepath.Join(bpDir, "fixtures", "windows_app"))
				app.Buildpacks = []string{extensionBuildpackName + "_buildpack", finalBuildpackName + "_buildpack"}
				app.Stack = os.Getenv("CF_STACK")
			})

			AfterEach(func() {
				Expect(cutlass.DeleteBuildpack(extensionBuildpackName)).To(Succeed())
				Expect(cutlass.DeleteBuildpack(finalBuildpackName)).To(Succeed())
				Expect(cutlass.RemovePackagedBuildpack(extensionBuildpack)).To(Succeed())
			})

			It("deploys successfully", func() {
				PushAppAndConfirm(app)
				if cutlass.Cached {
					Expect(app.Stdout.String()).ToNot(ContainSubstring("Download ["))

					// Expect "Copy [" twice
					Expect(app.Stdout.String()).To(MatchRegexp(`(?s)(?:Copy \[.*){2}`))
				} else {
					Expect(app.Stdout.String()).ToNot(ContainSubstring("Copy ["))

					// Expect "Download [" twice
					Expect(app.Stdout.String()).To(MatchRegexp(`(?s)(?:Download \[.*){2}`))
				}
				env, err := app.GetBody("/env")
				Expect(err).NotTo(HaveOccurred())
				Expect(env).To(ContainSubstring(`\\deps\\1\\bin`))
				Expect(env).To(ContainSubstring(`\\deps\\0\\bin`))
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
