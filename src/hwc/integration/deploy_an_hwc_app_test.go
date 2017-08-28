package integration_test

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CF HWC Buildpack", func() {
	var app *cutlass.App
	AfterEach(func() { app = DestroyApp(app) })

	Describe("deploying an hwc app", func() {
		BeforeEach(func() {
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
})
