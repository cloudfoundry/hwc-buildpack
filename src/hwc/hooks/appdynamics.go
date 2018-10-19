package hooks

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudfoundry/libbuildpack"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type AppdynamicsHook struct {
	libbuildpack.DefaultHook
	Log *libbuildpack.Logger
}

type Plan struct {
	Credentials Credential `json:"credentials"`
}

type Credential struct {
	ControllerHost   string `json:"host-name"`
	ControllerPort   string `json:"port"`
	SslEnabled       bool   `json:"ssl-enabled"`
	AccountAccessKey string `json:"account-access-key"`
	AccountName      string `json:"account-name"`
}

type VcapApplication struct {
	ApplicationName  string `json:"application_name"`
	ApplicationId    string `json:"application_id"`
	ApplicationSpace string `json:"space_name"`
}

func (h AppdynamicsHook) getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func (h AppdynamicsHook) CopyAppDynamicsAgentFromVendor(stager *libbuildpack.Stager) error {
	h.Log.BeginStep("Copying AppDynamics Agent files")
	return nil
}

func (h AppdynamicsHook) CreateEnv(controllerConfig Credential, applicationConfig VcapApplication) (map[string]string, error) {
	appdEnv := map[string]string{
		"COR_ENABLE_PROFILING":               "1",
		"COR_PROFILER":                       "{39AEABC1-56A5-405F-B8E7-C3668490DB4A}",
		"COR_PROFILER_PATH_64":               `%HOME\.appdynamics\AppDynamics.Profiler_x64.dll`,
		"appdynamics.controller.hostName":    controllerConfig.ControllerHost,
		"appdynamics.agent.accountAccessKey": controllerConfig.AccountAccessKey,
		"appdynamics.agent.accountName":      controllerConfig.AccountName,
		"appdynamics.controller.ssl.enabled": strconv.FormatBool(controllerConfig.SslEnabled),
		"APPDYNAMICS_AGENT_APPLICATION_NAME": h.getEnv("APPDYNAMICS_AGENT_APPLICATION_NAME", applicationConfig.ApplicationName),
		"APPDYNAMICS_AGENT_TIER_NAME":        h.getEnv("APPDYNAMICS_AGENT_TIER_NAME", applicationConfig.ApplicationSpace),
		"APPDYNAMICS_AGENT_NODE_NAME":        h.getEnv("APPDYNAMICS_AGENT_NODE_NAME", applicationConfig.ApplicationSpace),
	}
	return appdEnv, nil
}

func (h AppdynamicsHook) CreateDefaultEnv(applicationConfig VcapApplication) (map[string]string, error) {
	appdEnv := map[string]string{
		"COR_ENABLE_PROFILING": "1",
		"COR_PROFILER":         "{39AEABC1-56A5-405F-B8E7-C3668490DB4A}",
		"COR_PROFILER_PATH_64": `%HOME\.appdynamics\AppDynamics.Profiler_x64.dll`,
	}
	return appdEnv, nil
}

func (h AppdynamicsHook) ParseAppDynamicsVcapService() (Credential, error) {
	vcapServices := os.Getenv("VCAP_SERVICES")

	services := make(map[string][]Plan)
	if err := json.Unmarshal([]byte(vcapServices), &services); err != nil {
		h.Log.Debug("Could not unmarshal VCAP_SERVICES JSON exiting")
		return Credential{}, err
	}

	val, pres := services["appdynamics"]
	if !pres {
		h.Log.Debug("Not bound to AppDynamics")
		return Credential{}, errors.New("service appdynamics not present")
	}
	return val[0].Credentials, nil
}

func (h AppdynamicsHook) ParseVcapApplication() (VcapApplication, error) {
	vcapApplication := os.Getenv("VCAP_APPLICATION")
	application := VcapApplication{}
	if err := json.Unmarshal([]byte(vcapApplication), &application); err != nil {
		h.Log.Debug("Could not unmarshal VCAP_APPLICATION JSON")
		return VcapApplication{}, err
	}
	return application, nil
}

func (h AppdynamicsHook) WriteEnvFile(stager *libbuildpack.Stager) error {
	controllerConfig, err := h.ParseAppDynamicsVcapService()
	if err != nil {
		return err
	}

	applicationConfig, err := h.ParseVcapApplication()
	if err != nil {
		return err
	}

	var appdEnv map[string]string

	defaultJsonConfigFile := filepath.Join(stager.BuildDir(), ".appdynamics", "AppDynamicsConfig.json.default")
	jsonConfigFile := filepath.Join(stager.BuildDir(), ".appdynamics", "AppDynamicsConfig.json")

	if _, err := os.Stat(jsonConfigFile); os.IsNotExist(err) {

		h.Log.BeginStep("Writing AppDynamics Configuration")
		h.Log.BeginStep("Copying %v to %v", defaultJsonConfigFile, jsonConfigFile)
		if _, err = copy(defaultJsonConfigFile, jsonConfigFile); err != nil {
			return err
		}

		if appdEnv, err = h.CreateEnv(controllerConfig, applicationConfig); err != nil {
			return err
		}
	} else {
		h.Log.BeginStep("Using AppDynamicsConfig.json in .appdynamics directory")
		if appdEnv, err = h.CreateDefaultEnv(applicationConfig); err != nil {
			return err
		}
	}

	scriptContents := "echo [AppDynamics]Creating AppDynamics Environment"

	for envKey, envVal := range appdEnv {
		envStr := fmt.Sprintf("set %s=%s", envKey, envVal)
		scriptContents += "\n" + envStr
	}

	if err = stager.WriteProfileD("appd.bat", scriptContents); err != nil {
		return err
	}

	return nil
}

func (h AppdynamicsHook) BeforeCompile(stager *libbuildpack.Stager) error {
	h.Log.BeginStep("Setting up AppDynamics")

	if err := h.CopyAppDynamicsAgentFromVendor(stager); err != nil {
		h.Log.Debug("%v", err)
		return nil
	}

	if err := h.WriteEnvFile(stager); err != nil {
		h.Log.Debug("%v", err)
		return nil
	}

	return nil
}

func init() {
	logger := libbuildpack.NewLogger(os.Stdout)
	libbuildpack.AddHook(AppdynamicsHook{
		Log: logger,
	})
}
