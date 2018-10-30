package finalize_test

//go:generate mockgen -source=finalize.go --destination=mocks_test.go --package=finalize_test
import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"

	"github.com/cloudfoundry/hwc-buildpack/src/hwc/finalize"

	"github.com/cloudfoundry/libbuildpack"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Finalize", func() {
	var (
		err            error
		buildDir       string
		finalizer      finalize.Finalizer
		mockCtrl       *gomock.Controller
		mockManifest   *MockManifest
		mockStager     *MockStager
		mockCommand    *MockCommand
		mockHarmonizer *MockHarmonizer
	)

	BeforeEach(func() {
		buildDir, err = ioutil.TempDir("", "hwc-buildpack.build.")
		buffer := new(bytes.Buffer)
		logger := libbuildpack.NewLogger(buffer)

		mockCtrl = gomock.NewController(GinkgoT())
		mockManifest = NewMockManifest(mockCtrl)
		mockStager = NewMockStager(mockCtrl)
		mockCommand = NewMockCommand(mockCtrl)
		mockHarmonizer = NewMockHarmonizer(mockCtrl)

		finalizer = finalize.Finalizer{
			BuildDir:   buildDir,
			Manifest:   mockManifest,
			Stager:     mockStager,
			Command:    mockCommand,
			Harmonizer: mockHarmonizer,
			Log:        logger,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()

		err = os.RemoveAll(buildDir)
		Expect(err).To(BeNil())
	})

	Describe("Run", func() {
		Describe("success", func() {
			It("runs the hwc functions", func() {
				mockHarmonizer.EXPECT().CheckWebConfig().Return(nil)
				mockHarmonizer.EXPECT().LinkLegacyHwc().Return(nil)

				err = finalizer.Run()
				Expect(err).To(BeNil())
			})
		})

		Describe("errors", func() {
			It("runs the hwc functions", func() {
				mockHarmonizer.EXPECT().CheckWebConfig().Return(errors.New("BOOM"))

				err = finalizer.Run()
				Expect(err).To(HaveOccurred())
			})

			It("runs the hwc functions", func() {
				mockHarmonizer.EXPECT().CheckWebConfig().Return(nil)
				mockHarmonizer.EXPECT().LinkLegacyHwc().Return(errors.New("BOOM"))

				err = finalizer.Run()
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
