package supply_test

import (
	"bytes"
	"errors"
	"path/filepath"

	"github.com/cloudfoundry/hwc-buildpack/src/hwc/supply"

	"github.com/cloudfoundry/libbuildpack"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate mockgen -source=supply.go --destination=mocks_test.go --package=supply_test

var _ = Describe("Supply", func() {
	var (
		supplier      supply.Supplier
		mockCtrl      *gomock.Controller
		mockManifest  *MockManifest
		mockInstaller *MockInstaller
		mockStager    *MockStager
		mockCommand   *MockCommand
	)

	BeforeEach(func() {
		buffer := new(bytes.Buffer)
		logger := libbuildpack.NewLogger(buffer)

		mockCtrl = gomock.NewController(GinkgoT())
		mockManifest = NewMockManifest(mockCtrl)
		mockStager = NewMockStager(mockCtrl)
		mockCommand = NewMockCommand(mockCtrl)
		mockInstaller = NewMockInstaller(mockCtrl)

		supplier = supply.Supplier{
			Manifest:  mockManifest,
			Installer: mockInstaller,
			Stager:    mockStager,
			Command:   mockCommand,
			Log:       logger,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Run", func() {
		var (
			depDir      string
			expectedDep libbuildpack.Dependency
			expectedDir string
		)

		BeforeEach(func() {
			depDir = "some-dep-dir"
			mockStager.EXPECT().DepDir().Return(depDir)
			expectedDep = libbuildpack.Dependency{Name: "hwc", Version: "12.0.0"}
			mockManifest.EXPECT().DefaultVersion("hwc").Return(libbuildpack.Dependency{Name: "hwc", Version: "12.0.0"}, nil)
			expectedDir = filepath.Join(depDir, "hwc")
		})

		Context("the installer succeeds", func() {
			It("installs the hwc dependency to <dep-dir>/hwc", func() {
				mockInstaller.EXPECT().InstallDependency(expectedDep, expectedDir).Return(nil)

				Expect(supplier.Run()).To(Succeed())
			})
		})

		Context("the installer fails to install the dependency", func() {
			It("returns a helpful error", func() {
				mockInstaller.EXPECT().InstallDependency(expectedDep, expectedDir).Return(errors.New("some installer error"))

				err := supplier.Run()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("some installer error"))
			})
		})
	})
})
