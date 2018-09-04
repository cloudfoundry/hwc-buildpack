package brats_test

import (
	"github.com/cloudfoundry/libbuildpack/bratshelper"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("hwc buildpack", func() {
	// Unbuilt hwc-buildpack not yet supported
	//bratshelper.UnbuiltBuildpack("hwc", CopyBrats)

	// BRATs helper needs to change to support windows .profile.bat scripts
	//bratshelper.DeployAppWithExecutableProfileScript("hwc", CopyBrats)

	bratshelper.DeployingAnAppWithAnUpdatedVersionOfTheSameBuildpack(CopyBrats)
	bratshelper.DeployAnAppWithSensitiveEnvironmentVariables(CopyBrats)

	// bratshelper.ForAllSupportedVersions("hwc", CopyBrats, func(hwcVersion string, app *cutlass.App) {
	// 	bratshelper.PushApp(app)
	//
	// 	By("installs the correct version of hwc", func() {
	// 		Expect(app.Stdout.String()).To(ContainSubstring("Installing hwc " + hwcVersion))
	// 	})
	// 	By("runs a simple webserver", func() {
	// 		Expect(app.GetBody("/")).To(ContainSubstring("hello i am nora"))
	// 	})
	// })
})
