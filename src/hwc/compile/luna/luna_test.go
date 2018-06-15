package luna_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudfoundry/libbuildpack"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"hwc/compile/luna"
)

var _ = Describe("Luna", func() {

	var (
		buildDir              string
		logger                *libbuildpack.Logger
		lunaDir               string
		serverCert1           string
		serverCert2           string
		serverCert3           string
		clientCert            string
		clientPrivateKey      string
		VCAP_SERVICES         string
		VCAP_SERVICES_NO_LUNA string
		expectedCrystokiIni   string
	)

	BeforeEach(func() {

		buildDir, _ = ioutil.TempDir("", "hwc-buildpack.build.")
		buffer := new(bytes.Buffer)
		logger = libbuildpack.NewLogger(buffer)
		lunaDir = filepath.Join(buildDir, ".cloudfoundry", ".luna")

		expectedCrystokiIni = `[Luna]
PEDTimeout1=100000
PEDTimeout2=200000
CommandTimeoutPedSet=720000
KeypairGenTimeOut=2700000
CloningCommandTimeOut=300000
PEDTimeout3=10000
DefaultTimeOut=500000
[LunaSA Client]
ReceiveTimeout=20000
TCPKeepAlive=1
NetClient=1
ServerCAFile=.cloudfoundry\.luna\serverCertificates.pem
ClientCertFile=.cloudfoundry\.luna\clientCertificate.pem
ClientPrivKeyFile=.cloudfoundry\.luna\clientPrivateKey.pem
ServerName00=server1
ServerPort00=1792
ServerHtl00=0
ServerName01=server2
ServerPort01=1792
ServerHtl01=0
ServerName02=server3
ServerPort02=1792
ServerHtl02=0
[Misc]
PE1746Enabled=0
[HASynchronize]
par1=1
par2=1
[HAConfiguration]
HAOnly=1
reconnAtt=-1
AutoReconnectInterval=60
[VirtualToken]
VirtualToken00Label=par1
VirtualToken00SN=1888800001
VirtualToken00Members=888800001,888800002
VirtualToken01Label=par2
VirtualToken01SN=1888800003
VirtualToken01Members=888800003
`
		expectedCrystokiIni = strings.Replace(expectedCrystokiIni, "\n", "\r\n", -1)

		serverCert1 = `-----BEGIN CERTIFICATE-----
serverCert1
-----END CERTIFICATE-----
`
		serverCert2 = `-----BEGIN CERTIFICATE-----
serverCert2
-----END CERTIFICATE-----
`
		serverCert3 = `-----BEGIN CERTIFICATE-----
serverCert3
-----END CERTIFICATE-----
`
		clientCert = `-----BEGIN CERTIFICATE-----
clientCert
-----END CERTIFICATE-----
`
		clientPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
clientPrivateKey
-----END RSA PRIVATE KEY-----
`
		vcap := map[string]interface{}{
			"user-provided": []interface{}{
				map[string]interface{}{
					"credentials": map[string]interface{}{
						"client": map[string]interface{}{
							"certificate": clientCert,
							"private-key": clientPrivateKey,
						},
						"servers": []interface{}{
							map[string]interface{}{
								"name":        "server1",
								"certificate": serverCert1,
							},
							map[string]interface{}{
								"name":        "server2",
								"certificate": serverCert2,
							},
							map[string]interface{}{
								"name":        "server3",
								"certificate": serverCert3,
							},
						},
						"groups": []interface{}{
							map[string]interface{}{
								"label": "par1",
								"members": []interface{}{
									"888800001", "888800002",
								},
							},
							map[string]interface{}{
								"label": "par2",
								"members": []interface{}{
									"888800003",
								},
							},
						},
					},
					"label": "user-provided",
					"name":  "myluna",
				},
			},
		}
		vcap_byte, _ := json.Marshal(vcap)
		VCAP_SERVICES = string(vcap_byte[:])

		vcap = map[string]interface{}{
			"user-provided": []interface{}{
				map[string]interface{}{
					"credentials": map[string]interface{}{
						"secret": "open sesame",
					},
					"label": "user-provided",
					"name":  "notluna",
				},
			},
		}
		vcap_byte, _ = json.Marshal(vcap)
		VCAP_SERVICES_NO_LUNA = string(vcap_byte[:])
	})

	AfterEach(func() {
		err := os.RemoveAll(buildDir)
		Expect(err).To(BeNil())
	})

	Describe("Test Install Luna", func() {
		Context("Should configure luna client", func() {
			It("Should make .luna directory and write ntls credentials and crytoki.ini files", func() {
				os.Setenv("VCAP_SERVICES", VCAP_SERVICES)
				luna := luna.NewLuna(logger, buildDir)
				err := luna.InstallLuna()
				Expect(err).To(BeNil())
				clientCertActual, err := ioutil.ReadFile(filepath.Join(lunaDir, "clientCertificate.pem"))
				Expect(err).To(BeNil())
				Expect(string(clientCertActual)).To(Equal(clientCert))
				clientPrivateKeyActual, err := ioutil.ReadFile(filepath.Join(lunaDir, "clientPrivateKey.pem"))
				Expect(err).To(BeNil())
				Expect(string(clientPrivateKeyActual)).To(Equal(clientPrivateKey))
				serverCertsActual, err := ioutil.ReadFile(filepath.Join(lunaDir, "serverCertificates.pem"))
				Expect(err).To(BeNil())
				Expect(string(serverCertsActual)).To(Equal(serverCert1 + serverCert2 + serverCert3))
				crystokiIniActual, err := ioutil.ReadFile(filepath.Join(lunaDir, "crystoki.ini"))
				Expect(err).To(BeNil())
				Expect(string(crystokiIniActual)).To(Equal(expectedCrystokiIni))
			})
		})
	})

	Context("Should not configure luna client", func() {
		It("Should not create .luna directory", func() {
			os.Setenv("VCAP_SERVICES", VCAP_SERVICES_NO_LUNA)
			luna := luna.NewLuna(logger, lunaDir)
			err := luna.InstallLuna()
			Expect(err).To(BeNil())
			_, err = os.Stat(lunaDir)
			Expect(os.IsNotExist(err)).To(Equal(true))
		})
	})
})
