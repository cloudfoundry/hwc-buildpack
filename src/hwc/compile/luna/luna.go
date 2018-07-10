package luna

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/cloudfoundry/libbuildpack"
)

type Luna struct {
	Log                    *libbuildpack.Logger
	BuildDir               string
	LunaDir                string
	credentials            *LunaCredentials
	ClientPrivateKeyPath   string
	ClientCertificatePath  string
	ServerCertificatesPath string
}

type Client struct {
	Certificate string
	PrivateKey  string `json:"private-key"`
}

type Server struct {
	Name        string
	Certificate string
}

type Group struct {
	Label   string
	Members []string
}

type LunaCredentials struct {
	Client  Client
	Servers []Server
	Groups  []Group
}

func NewLuna(log *libbuildpack.Logger, buildDir string) *Luna {
	l := new(Luna)
	l.Log = log
	l.BuildDir = buildDir
	l.LunaDir = filepath.Join(l.BuildDir, ".cloudfoundry", ".luna")
	l.ClientPrivateKeyPath = filepath.Join(l.LunaDir, "clientPrivateKey.pem")
	l.ClientCertificatePath = filepath.Join(l.LunaDir, "clientCertificate.pem")
	l.ServerCertificatesPath = filepath.Join(l.LunaDir, "serverCertificates.pem")
	return l
}

func validateCredentials(credentials *LunaCredentials) bool {
	if credentials.Client.Certificate == "" || credentials.Client.PrivateKey == "" {
		return false
	}
	if len(credentials.Servers) == 0 {
		return false
	}
	for _, server := range credentials.Servers {
		if server.Name == "" || server.Certificate == "" {
			return false
		}
	}
	for _, group := range credentials.Groups {
		if group.Label == "" || len(group.Members) == 0 {
			return false
		}
	}
	return true
}

func checkServiceName(service map[string]interface{}) bool {
	nameInterface, ok := service["name"]
	if !ok {
		return false
	}
	name, ok := nameInterface.(string)
	if !ok {
		return false
	}
	name = strings.ToLower(name)
	return strings.Contains(name, "luna") || strings.Contains(name, "hsm")
}

func (l *Luna) parseVCAPServices() {
	vcapServices := os.Getenv("VCAP_SERVICES")
	var f map[string]interface{}
	json.Unmarshal([]byte(vcapServices), &f)
	for key, _ := range f {
		services, ok := f[key].([]interface{})
		if !ok {
			continue
		}
		for _, v1 := range services {

			service, ok := v1.(map[string]interface{})
			if !ok {
				continue
			}

			if !checkServiceName(service) {
				continue
			}

			credsInterface, ok := service["credentials"]
			if !ok {
				continue
			}
			credentialsJson, err := json.Marshal(credsInterface)
			if err != nil {
				continue
			}
			var credentials LunaCredentials
			err = json.Unmarshal(credentialsJson, &credentials)
			if err == nil && validateCredentials(&credentials) {
				l.credentials = &credentials
				return
			}
		}
	}
}

func (l *Luna) createLunaDirectory() error {
	return os.MkdirAll(l.LunaDir, 0755)
}

func (l *Luna) writeNTLSCredentials() error {

	privateKey := []byte(l.credentials.Client.PrivateKey)
	err := ioutil.WriteFile(l.ClientPrivateKeyPath, privateKey, 0640)
	if err != nil {
		return err
	}

	clientCertificate := []byte(l.credentials.Client.Certificate)
	err = ioutil.WriteFile(l.ClientCertificatePath, clientCertificate, 0644)
	if err != nil {
		return err
	}

	var serverCertificates string
	serverCertificates = ""
	for _, server := range l.credentials.Servers {
		serverCertificates += server.Certificate
	}
	err = ioutil.WriteFile(l.ServerCertificatesPath, []byte(serverCertificates), 0644)

	return err
}

type Section struct {
	Name     string
	Settings []Setting
}

type Setting struct {
	Name  string
	Value string
}

func (l *Luna) writeCrystokiIni() error {

	sections := []Section{}

	luna := []Setting{{"PEDTimeout1", "100000"}, {"PEDTimeout2", "200000"},
		{"CommandTimeoutPedSet", "720000"}, {"KeypairGenTimeOut", "2700000"}, {"CloningCommandTimeOut", "300000"},
		{"PEDTimeout3", "10000"}, {"DefaultTimeOut", "500000"}}

	sections = append(sections, Section{"Luna", luna})

	lunaSAClient := []Setting{{"ReceiveTimeout", "20000"}, {"TCPKeepAlive", "1"}, {"NetClient", "1"},
		{"ServerCAFile", ".cloudfoundry\\.luna\\serverCertificates.pem"}, {"ClientCertFile", ".cloudfoundry\\.luna\\clientCertificate.pem"},
		{"ClientPrivKeyFile", ".cloudfoundry\\.luna\\clientPrivateKey.pem"}}

	for index, server := range l.credentials.Servers {
		lunaSAClient = append(lunaSAClient, Setting{fmt.Sprintf("ServerName%02d", index), server.Name})
		lunaSAClient = append(lunaSAClient, Setting{fmt.Sprintf("ServerPort%02d", index), "1792"})
		lunaSAClient = append(lunaSAClient, Setting{fmt.Sprintf("ServerHtl%02d", index), "0"})
	}

	sections = append(sections, Section{"LunaSA Client", lunaSAClient})

	sections = append(sections, Section{"Misc", []Setting{{"PE1746Enabled", "0"}}})

	haSynchronize := []Setting{}
	for _, group := range l.credentials.Groups {
		haSynchronize = append(haSynchronize, Setting{group.Label, "1"})
	}
	sections = append(sections, Section{"HASynchronize", haSynchronize})

	sections = append(sections, Section{"HAConfiguration", []Setting{{"HAOnly", "1"}, {"reconnAtt", "-1"}, {"AutoReconnectInterval", "60"}}})

	virtualToken := []Setting{}
	for index, group := range l.credentials.Groups {
		virtualToken = append(virtualToken, Setting{fmt.Sprintf("VirtualToken%02dLabel", index), group.Label})
		virtualToken = append(virtualToken, Setting{fmt.Sprintf("VirtualToken%02dSN", index), "1" + group.Members[0]})
		virtualToken = append(virtualToken, Setting{fmt.Sprintf("VirtualToken%02dMembers", index), strings.Join(group.Members, ",")})
	}
	sections = append(sections, Section{"VirtualToken", virtualToken})

	f, err := os.Create(filepath.Join(l.LunaDir, "crystoki.ini"))
	if err != nil {
		return err
	}
	defer f.Close()

	var iniTemplate = "{{range $index, $section := .}}[{{.Name}}]\r\n{{range $index2, $setting := $section.Settings}}{{.Name}}={{.Value}}\r\n{{end}}{{end}}"
	t := template.New("ini")
	t, err = t.Parse(iniTemplate)
	if err != nil {
		return err
	}
	return t.Execute(f, sections)
}

func (l *Luna) writeProfileD() error {
	profileD := filepath.Join(l.BuildDir, ".profile.d")
	err := os.MkdirAll(profileD, 0755)
	if err != nil {
		return err
	}
	lunaBatch := "set ChrystokiConfigurationPath=.cloudfoundry\\.luna\r\n"
	return ioutil.WriteFile(filepath.Join(profileD, "000_luna.bat"), []byte(lunaBatch), 0644)
}

func (l *Luna) InstallLuna() error {
	l.parseVCAPServices()
	if l.credentials == nil {
		l.Log.Info("No Luna Service found: not installing Luna")
		return nil
	}
	l.Log.Info("Installing Luna to %s", l.LunaDir)
	err := l.createLunaDirectory()
	if err != nil {
		return err
	}
	err = l.writeNTLSCredentials()
	if err != nil {
		return err
	}
	err = l.writeCrystokiIni()
	if err != nil {
		return err
	}
	err = l.writeProfileD()
	if err != nil {
		return err
	}
	l.Log.Info("Completed installing Luna")

	return nil
}
