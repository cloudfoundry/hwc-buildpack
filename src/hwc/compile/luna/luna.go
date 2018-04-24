package luna

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudfoundry/libbuildpack"
)

type Luna struct {
	Log                    *libbuildpack.Logger
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

func NewLuna(log *libbuildpack.Logger, lunaDir string) *Luna {
	l := new(Luna)
	l.Log = log
	l.LunaDir = lunaDir
	l.ClientPrivateKeyPath = filepath.Join(lunaDir, "clientPrivateKey.pem")
	l.ClientCertificatePath = filepath.Join(lunaDir, "clientCertificate.pem")
	l.ServerCertificatesPath = filepath.Join(lunaDir, "serverCertificates.pem")
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
			creds1, ok := v1.(map[string]interface{})
			if !ok {
				continue
			}
			creds2, ok := creds1["credentials"]
			if !ok {
				continue
			}
			credentialsJson, err := json.Marshal(creds2)
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

func (l *Luna) writeCrystokiIni() error {
	var conf string
	conf = ""
	lunaSection := `[Luna]
PEDTimeout1=100000
PEDTimeout2=200000
CommandTimeoutPedSet=720000
KeypairGenTimeOut=2700000
CloningCommandTimeOut=300000
PEDTimeout3=10000
DefaultTimeOut=500000
`

	conf += lunaSection

	client := `[LunaSA Client]
ReceiveTimeout=20000
TCPKeepAlive=1
NetClient=1
`
	client += "ServerCAFile=.cloudfoundry\\.luna\\serverCertificates.pem\n"
	client += "ClientCertFile=.cloudfoundry\\.luna\\clientCertificate.pem\n"
	client += "ClientPrivKeyFile=.cloudfoundry\\.luna\\clientPrivateKey.pem\n"

	for index, server := range l.credentials.Servers {
		client += fmt.Sprintf("ServerName%02d=", index) + server.Name + "\n"
		client += fmt.Sprintf("ServerPort%02d=1792", index) + "\n"
		client += fmt.Sprintf("ServerHtl%02d=0", index) + "\n"
	}

	conf += client

	misc := `[Misc]
PE1746Enabled=0
`

	conf += misc

	conf += "[HASynchronize]\n"
	for _, group := range l.credentials.Groups {
		conf += group.Label + "=1\n"
	}

	haconf := `[HAConfiguration]
HAOnly=1
reconnAtt=-1
AutoReconnectInterval=60
`
	conf += haconf
	conf += "[VirtualToken]\n"
	for index, group := range l.credentials.Groups {
		conf += fmt.Sprintf("VirtualToken%02dLabel=%s\n", index, group.Label)
		conf += fmt.Sprintf("VirtualToken%02dSN=1%s\n", index, group.Members[0])
		members := strings.Join(group.Members, ",")
		conf += fmt.Sprintf("VirtualToken%02dMembers=%s\n", index, members)
	}

	filepath.Join(l.LunaDir, "crystoki.ini")
	conf = strings.Replace(conf, "\n", "\r\n", -1)
	err := ioutil.WriteFile(filepath.Join(l.LunaDir, "crystoki.ini"), []byte(conf), 0666)
	if err != nil {
		return err
	}
	return err
}

func (l *Luna) InstallLuna() error {
	l.Log.Info("Installing luna to %s", l.LunaDir)
	l.parseVCAPServices()
	if l.credentials == nil {
		l.Log.Info("No Luna Service found in VCAP_SERVICES")
		return nil
	}
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
	l.Log.Info("Completed installing Luna")

	return nil
}
