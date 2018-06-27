package compile

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
	"crypto/sha256"
	"encoding/hex"

	"github.com/cloudfoundry/libbuildpack"
)

const bucketXMLUrl = "https://nr-downloads-main.s3.amazonaws.com/?delimiter=/&prefix=dot_net_agent/latest_release/"
const nrVersionPattern = "((\\d{1,3}\\.){3}\\d{1,3})"
const latestNrDownloadUrl = "http://download.newrelic.com/dot_net_agent/latest_release/newrelic-agent-win-x64-9.9.9.9.zip"
const latestNrDownloadSha256Url = "http://download.newrelic.com/dot_net_agent/latest_release/SHA256/newrelic-agent-win-x64-9.9.9.9.zip.sha256"

type bucketResultXMLNode struct {
	XMLName xml.Name
	Content []byte                `xml:",innerxml"`
	Nodes   []bucketResultXMLNode `xml:",any"`
}

// RULES for installing newrelic agent:
//	if:
//		- NEW_RELIC_LICENSE_KEY exists
//		- NEW_RELIC_DOWNLOAD_URL exists
//		- there is a user-provided-service with the word "newrelic" in the name
//		- there is a SERVICE in VCAP_SERVICES with the name "newrelic"
//		- for cached buildpack: nrDownloadFile is set to file name (non-blank)
//	then call c.InstallNewRelic()

func (c *Compiler) InstallNewRelic() error {

	buildpackDir := getBuildpackDir(c)

	// check if the app requires to bind to new relic agent
	bindNrAgent := false
	if _, exists := os.LookupEnv("NEW_RELIC_LICENSE_KEY"); exists {
		bindNrAgent = true
	} else if _, exists := os.LookupEnv("NEW_RELIC_DOWNLOAD_URL"); exists {
		bindNrAgent = true
	} else {
		vCapServicesEnvValue := os.Getenv("VCAP_SERVICES")
		if vCapServicesEnvValue != "" {
			var vcapServices map[string]interface{}
			if err := json.Unmarshal([]byte(vCapServicesEnvValue), &vcapServices); err != nil {
		    	c.Log.Error("", err)
			} else {
		    	// check for a service from newrelic service broker (or tile)
				if _, exists := vcapServices["newrelic"].([]interface{}); exists {
					bindNrAgent = true
				} else {
			    	// check user-provided-services
					userProvidesServicesElement, _ := vcapServices["user-provided"].([]interface{})
			        for _, ups := range userProvidesServicesElement {
			        	s, _ := ups.(map[string]interface{})
			        	if exists := strings.Contains(strings.ToLower(s["name"].(string)), "newrelic"); exists {
			        		bindNrAgent = true
			        		break; 
						}
					}
				}
			}
		}
	}
	if !bindNrAgent { // there is no references in the environment to new relic agent
		return nil
	}
	// ############################################################################################
	// ############################################################################################

	c.Log.BeginStep("Installing NewRelic .Net Framework Agent")

	nrDownloadURL := "http://download.newrelic.com/dot_net_agent/latest_release/newrelic-agent-win-x64-0.0.0.0.zip"
	nrDownloadFile := ""
	nrVersion := "latest"
	nrSha256Sum := ""
	for _, entry := range c.Manifest.(*libbuildpack.Manifest).ManifestEntries {
		if entry.Dependency.Name == "newrelic" {
			nrDownloadURL = entry.URI
			nrVersion = entry.Dependency.Version
			nrDownloadFile = entry.File
			nrSha256Sum = entry.SHA256
		}
	}

	// ############################################################################################
	// ############################################################################################

	newrelicDir := filepath.Join(c.BuildDir, "newrelic")

	c.Log.Info("Creating tmp folder for downloading agent")
	tmpDir, err := ioutil.TempDir(c.BuildDir, "downloads")
	if err != nil {
		return err
	}

	nrDownloadLocalFilename := filepath.Join(tmpDir, "NewRelic.Agent.Installer.zip")

	needToDownloadNrAgdentFile := false

	if downloadURLOverride, exists := os.LookupEnv("NEW_RELIC_DOWNLOAD_URL"); exists == true {
		nrDownloadURL = strings.TrimSpace(downloadURLOverride)
		nrSha256Sum = "" // when NEW_RELIC_DOWNLOAD_URL is used ignore sha256 sum
		needToDownloadNrAgdentFile = true
	} else if nrDownloadFile != "" { // this file is cached by the buildpack
		c.Log.Info("Using cached dependencies...")
		source := nrDownloadFile
		if !filepath.IsAbs(source) {
			source = filepath.Join(buildpackDir, source)
		}
		c.Log.Info("Copy [%s]", source)
		if err := libbuildpack.CopyFile(source, nrDownloadLocalFilename); err != nil {
			return err
		}
	} else {
		if (nrDownloadURL == "" || in_array(strings.ToLower(nrVersion), []string{"", "0.0.0.0", "latest", "current"})) {

			c.Log.Info("Evaluating latest agent version ")
			nrLatestAgentVersion, err := evalLatestAgentVersion()
			if err != nil {
				c.Log.Error("Unable to evaluate agent version from the metadata bucket", err)
				return err
			}
			nrVersion = nrLatestAgentVersion
			c.Log.Info("Latest New Relic agent version: %v", nrVersion)


			updatedUrl, err := substituteUrlVersion(latestNrDownloadUrl, nrVersion, c.Log)
			if err != nil {
				c.Log.Error("filed to substitute agent version in url")
				return err
			}
			nrDownloadURL = updatedUrl

			// read sha256 sum of the agent from NR download site
			latestNrAgentSha256Sum, err := getLatestNrAgentSha256Sum(tmpDir, nrVersion, c.Log)
			if err != nil {
				c.Log.Info("Can't get SHA256 checksum for latest New Relic Agent download")
				return err
			}
			nrSha256Sum = latestNrAgentSha256Sum
		} 

		c.Log.Info("Using New Relic version %s", nrVersion)
		needToDownloadNrAgdentFile = true
	}
	if needToDownloadNrAgdentFile {
		if err := downloadDependency(nrDownloadURL, nrDownloadLocalFilename, c.Log); err != nil {
			return err
		}
	}

	// compare sha256 sum of the downloaded file against expected sum
	if nrSha256Sum != "" {
		if err := checkSha256(nrDownloadLocalFilename, nrSha256Sum); err != nil {
			c.Log.Info("SHA256 checksum failed")
			return err
		}
	}

	c.Log.BeginStep("Extracting NewRelic .Net Agent to %s", nrDownloadLocalFilename)
	if err := libbuildpack.ExtractZip(nrDownloadLocalFilename, newrelicDir); err != nil {
		c.Log.Error("Error Extracting NewRelic .Net Agent", err)
		return err
	}

	// decide which newrelic.config file to use
	if err := getNewRelicConfigFile(newrelicDir, buildpackDir, c); err != nil {
		return err
	}
	// ############################################################################################
	// ############################################################################################

	if err := getProcfile(buildpackDir, c); err != nil {
		return err
	}

	newrelicAppName := parseVcapApplicationEnv(c)

	newrelicLicenseKey := ""
	// NEW_RELIC_LICENSE_KEY env var always overwrites other license keys
	if _, exists := os.LookupEnv("NEW_RELIC_LICENSE_KEY"); exists == false {
		vCapServicesEnvValue := os.Getenv("VCAP_SERVICES")
		if vCapServicesEnvValue != "" {
			var vcapServices map[string]interface{}
			if err := json.Unmarshal([]byte(vCapServicesEnvValue), &vcapServices); err != nil {
		    	c.Log.Error("", err)
			} else {
				newrelicLicenseKey = parseNewRelicService(vcapServices, c)
				appName, licenseKey := parseUserProvidedServices(vcapServices, newrelicAppName, newrelicLicenseKey, c)
				newrelicAppName = appName
				newrelicLicenseKey = licenseKey
			}
		}
	}

	if err := buildRunCmd(c.BuildDir, newrelicAppName, newrelicLicenseKey, c); err != nil {
		return err
	}

	// remove tmp dir
	err = os.RemoveAll(tmpDir)
	if err != nil {
		c.Log.Info("tmp folder not removed")
	}

	return nil
}

func in_array(searchStr string, array []string) bool {
    for _, v := range array {
        if  v == searchStr { // item found in array of strings
            return true
        }   
    }
    return false
}

func getBuildpackDir(c *Compiler) string {
	// get the buildpack directory
	buildpackDir, err := libbuildpack.GetBuildpackDir()
	if err != nil {
		c.Log.Error("Unable to determine buildpack directory: %s", err.Error())
	}
	return buildpackDir
}

func getProcfile(buildpackDir string, c *Compiler) error {
	procFileBundledWithApp := filepath.Join(c.BuildDir, "Procfile")
	procFileBundledWithAppExists, err := libbuildpack.FileExists(procFileBundledWithApp)
	if err != nil {
		// no Procfile found in the app folder
		procFileBundledWithAppExists = false
	}
	if procFileBundledWithAppExists {
		// Procfile exists in app folder
		c.Log.Info("Using Procfile provided in the app folder")
	} else {
		c.Log.Info("No Procfile found in the app folder")
		// looking for Procfile in the buildpack dir
		procFileBundledWithBuildPack := filepath.Join(buildpackDir, "Procfile")
		procFileDest := filepath.Join(c.BuildDir, "Procfile")
		procFileBundledWithBuildPackExists, err := libbuildpack.FileExists(procFileBundledWithBuildPack)
		if err != nil {
			c.Log.Error("Error checking if Procfile exists in buildpack", err)
			return err
		}
		if procFileBundledWithBuildPackExists {
			// Procfile exists in buidpack folder
			c.Log.Info("Using Procfile provided with the buildpack")
			if err := libbuildpack.CopyFile(procFileBundledWithBuildPack, procFileDest); err != nil {
				c.Log.Error("Error copying Procfile provided by the buildpack", err)
				return err
			}
			c.Log.Info("Copied Procfile")
		} else {
			c.Log.Info("No Procfile provided by the buildpack")
		}
	}
	return nil
}

func getNewRelicConfigFile(newrelicDir string, buildpackDir string, c *Compiler) error {
	newrelicConfigBundledWithApp := filepath.Join(c.BuildDir, "newrelic.config")
	newrelicConfigDest := filepath.Join(newrelicDir, "newrelic.config")
	newrelicConfigBundledWithAppExists, err := libbuildpack.FileExists(newrelicConfigBundledWithApp)
	if err != nil {
		c.Log.Error("Unable to test existence of newrelic.config in app folder", err)
		newrelicConfigBundledWithAppExists = false
	}
	if newrelicConfigBundledWithAppExists {
		// newrelic.config exists in app folder
		c.Log.Info("Overwriting newrelic.config provided with app")
		if err := libbuildpack.CopyFile(newrelicConfigBundledWithApp, newrelicConfigDest); err != nil {
			c.Log.Error("Error Copying newrelic.config provided within the app folder", err)
			return err
		}
	} else {
		// check if newrelic.config exists in the buildpack folder
		newrelicConfigBundledWithBuildPack := filepath.Join(buildpackDir, "newrelic.config")
		newrelicConfigFileExists, err := libbuildpack.FileExists(newrelicConfigBundledWithBuildPack)
		if err != nil {
			c.Log.Error("Error checking if newrelic.confg exists in buildpack", err)
			return err
		}
		if newrelicConfigFileExists {
			// newrelic.config exists in buidpack folder
			c.Log.Info("Using newrelic.config provided with the buildpack")
			if err := libbuildpack.CopyFile(newrelicConfigBundledWithBuildPack, newrelicConfigDest); err != nil {
				c.Log.Error("Error copying newrelic.config provided by the buildpack", err)
				return err
			}
			c.Log.Info("Overwriting newrelic.config template provided with the buildpack")
		} else {
			c.Log.Info("Using default newrelic.config downloaded with the agent")
		}
	}
	return nil
}

func parseVcapApplicationEnv(c *Compiler) string {
	newrelicAppName := ""
	// NEW_RELIC_APP_NAME env var always overwrites other app names
	if _, exists := os.LookupEnv("NEW_RELIC_APP_NAME"); exists == false {
		vCapApplicationEnvValue := os.Getenv("VCAP_APPLICATION")
		var vcapApplication map[string]interface{}
		if err := json.Unmarshal([]byte(vCapApplicationEnvValue), &vcapApplication); err != nil {
			c.Log.Info("Unable to unmarshall VCAP_APPLICATION environment variable, NEW_RELIC_APP_NAME will not be set in profile script")
		} else {
			appName, ok := vcapApplication["application_name"].(string)
			if ok {
				c.Log.Info("VCAP_APPLICATION.application_name=" + appName)
				newrelicAppName = appName
			}
		}
	} else {
		newrelicAppName = ""
	}
	return newrelicAppName
}

func parseNewRelicService(vcapServices map[string]interface{}, c *Compiler) string {
	newrelicLicenseKey := ""
	// check for a service from newrelic service broker (or tile)
	newrelicElement, ok := vcapServices["newrelic"].([]interface{})
	if ok {
  		if len(newrelicElement) > 0 {
    		newrelicMap, ok := newrelicElement[0].(map[string]interface{})
    		if ok {
      			credMap, ok := newrelicMap["credentials"].(map[string]interface{})
      			if ok {
        			newrelicLicense, ok := credMap["licenseKey"].(string)
        			if ok {
          				c.Log.Info("VCAP_SERVICES.newrelic.credentials.licenseKey=" + newrelicLicense)
          				newrelicLicenseKey = newrelicLicense
        			}
      			}
    		}
  		}
	}
	return newrelicLicenseKey
}

func parseUserProvidedServices(vcapServices map[string]interface{}, newrelicAppName string, newrelicLicenseKey string, c *Compiler) (string, string) {
	// check user-provided-services
	userProvidesServicesElement, _ := vcapServices["user-provided"].([]interface{})
    for _, ups := range userProvidesServicesElement {
    	s, _ := ups.(map[string]interface{})
    	if found := strings.Contains(strings.ToLower(s["name"].(string)), "newrelic"); found == true {
			cmap, _ := s["credentials"].(map[string]interface{})
        	for key, cred := range cmap {
        		if (in_array(strings.ToLower(key), []string{"license_key", "licensekey", "new_relic_license_key"})) {
        			newrelicLicenseKey = cred.(string) // license key from user-provided-service -- overwrites license key from service broker
					c.Log.Info("VCAP_SERVICES." + s["name"].(string) + ".credentials." + key + "=" + newrelicLicenseKey)
				} else if (in_array(strings.ToLower(key), []string{"appname", "app_name", "new_relic_app_name"})) {
					newrelicAppName = cred.(string) // application name from user-provided-service -- overwrites name from service broker
					c.Log.Info("VCAP_SERVICES." + s["name"].(string) + ".credentials." + key + "=" + newrelicAppName)
				}
			}
		}
	}
	return newrelicAppName, newrelicLicenseKey
}

func buildRunCmd(buildDir string, newrelicAppName string, newrelicLicenseKey string, c *Compiler) error {
	runCmdFileDest := filepath.Join(c.BuildDir, "run.cmd")
	var scriptContentBuffer bytes.Buffer

	//Only the %~dp0 works. Using the newrelicDir wont work here
	dpSymbolNewRelic := `%~dp0newrelic`
	
	profilerSettings := "set COR_ENABLE_PROFILING=1\nset COR_PROFILER={71DA0A04-7777-4EC6-9643-7D28B46A8A41}\n"

	scriptContentBuffer.WriteString("set NEWRELIC_HOME=")
	scriptContentBuffer.WriteString(dpSymbolNewRelic)
	scriptContentBuffer.WriteString("\n")

	scriptContentBuffer.WriteString("set COR_PROFILER_PATH=")
	scriptContentBuffer.WriteString(filepath.Join(dpSymbolNewRelic, "NewRelic.Profiler.dll"))
	scriptContentBuffer.WriteString("\n")

	scriptContentBuffer.WriteString(profilerSettings)
	scriptContentBuffer.WriteString("\n")

	scriptContentBuffer.WriteString("set NEWRELIC_INSTALL_PATH=")
	scriptContentBuffer.WriteString(dpSymbolNewRelic)
	scriptContentBuffer.WriteString("\n")

	if newrelicAppName != "" {
		scriptContentBuffer.WriteString("set NEW_RELIC_APP_NAME=")
		scriptContentBuffer.WriteString(newrelicAppName)
		scriptContentBuffer.WriteString("\n")
	}

	if newrelicLicenseKey != "" {
		scriptContentBuffer.WriteString("set NEW_RELIC_LICENSE_KEY=")
		scriptContentBuffer.WriteString(newrelicLicenseKey)
		scriptContentBuffer.WriteString("\n")
	}

	scriptContentBuffer.WriteString(filepath.Join(".cloudfoundry", "hwc.exe"))
	scriptContentBuffer.WriteString("\n")

	scriptContents := scriptContentBuffer.String()

	err := writeToFile(strings.NewReader(scriptContents), runCmdFileDest, 0755)
	if err != nil {
		c.Log.Error("Unable to write run.cmd")
		return err
	}
	c.Log.Info("run.cmd file created to start hwc.exe with New Relic profiler variables")
	return nil
}

func writeToFile(source io.Reader, destFile string, mode os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(destFile), 0755)
	if err != nil {
		return err
	}

	fh, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer fh.Close()

	_, err = io.Copy(fh, source)
	if err != nil {
		return err
	}

	return nil
}

func checkSha256(filePath, expectedSha256 string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	sum := sha256.Sum256(content)

	actualSha256 := hex.EncodeToString(sum[:])

	if strings.ToLower(actualSha256) != strings.ToLower(expectedSha256) {
		return fmt.Errorf("dependency sha256 mismatch: expected sha256 %s, actual sha256 %s", expectedSha256, actualSha256)
	}
	return nil
}

func substituteUrlVersion(url string, nrVersion string, log *libbuildpack.Logger) (string, error) {
	nrVersionPatternMatcher, err := regexp.Compile(nrVersionPattern)
	if err != nil {
		log.Error("filed to build rexexp pattern matcher")
		return "", err
	}
	result := nrVersionPatternMatcher.FindStringSubmatch(url)
	if (len(result) <= 0) {
		return "", errors.New("Error: no version match found in url")
	}
	uriVersion := result[1] // version pattern found in the url

	return strings.Replace(url, uriVersion, nrVersion, -1), nil
}

func getLatestNrAgentSha256Sum(tmpDir string, latestNrVersion string, log *libbuildpack.Logger) (string, error) {
	shaUrl, err := substituteUrlVersion(latestNrDownloadSha256Url, latestNrVersion, log)
	if err != nil {
		log.Error("filed to substitute agent version in sha256 url")
		return "", err
	}

	sha256File := filepath.Join(tmpDir, "nragent.sha256")
	if err := downloadDependency(shaUrl, sha256File, log); err != nil {
		return "", err
	}

	sha256Sum, err := ioutil.ReadFile(sha256File)
	if err != nil {
		return "", err
	}

	return string(sha256Sum), nil
}

func downloadDependency(url string, filepath string, log *libbuildpack.Logger) (err error) {
	log.Info("Downloading from [%s]", url)
	log.Info("Saving to [%s]", filepath)

	var httpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func evalLatestAgentVersion() (string, error) {
	latestAgentVersion := ""
	resp, err := http.Get(bucketXMLUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Bad http status when downloading XML meta data: %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(data)
	dec := xml.NewDecoder(buf)

	var listBucketResultNode bucketResultXMLNode
	err = dec.Decode(&listBucketResultNode)
	if err != nil {
		return "", err
	}

	for _, nc := range listBucketResultNode.Nodes {
		if nc.XMLName.Local == "Contents" {
			key := ""
			for _, nc2 := range nc.Nodes {
				if nc2.XMLName.Local == "Key" {
					key = string(nc2.Content)
				}
			}
			nrVersionPatternMatcher, err := regexp.Compile(nrVersionPattern)
			if err != nil {
				return "", err
			}

			result := nrVersionPatternMatcher.FindStringSubmatch(key)
			if len(result) > 1 {
				latestAgentVersion = result[1]
			}
		}
	}
	return latestAgentVersion, nil
}
