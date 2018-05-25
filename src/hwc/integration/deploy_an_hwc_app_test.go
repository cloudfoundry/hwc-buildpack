package integration_test

import (
	"fmt"
	"os"
	"path/filepath"

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
	})
})
