package configparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	serverURLs                    []string
	balancingAlgorithmName        string
	serverTimeoutSeconds          int
	failedHealthChecksTillTimeout int
	slowStart                     bool
	slowStartSeconds              int
	stickySession                 bool
}

func (this *Config) InitDefault() {
	this.serverURLs = make([]string, 0)
	this.balancingAlgorithmName = "round_robin"
	this.serverTimeoutSeconds = 30
	this.failedHealthChecksTillTimeout = 3
	this.slowStart = false
	this.slowStartSeconds = 120
	this.stickySession = false
}

func (this *Config) InitFromFile(configFilePath string) {
	if configFilePath == "" {
		panic("Config file path can not be empty")
	}

	if _, err := os.Stat(configFilePath); err != nil {
		panic(err.Error())
	}

	this.InitDefault()
	this.loadConfigFromFile(configFilePath)
}

func (this *Config) Print() {
	fmt.Println(this)
}

func (this *Config) loadConfigFromFile(configFilePath string) {
	configDataJson := extractDataFromFile(configFilePath)
	data := unmarshalData(configDataJson)
	this.populateConfigWithExtractedData(data)
}

func extractDataFromFile(configFilePath string) []byte {
	configData, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		panic("Could not open config file\n" + err.Error())
	}

	return configData
}

func unmarshalData(configDataJson []byte) map[string]interface{} {
	var data map[string]interface{}

	err := json.Unmarshal(configDataJson, &data)

	if err != nil {
		panic("Invalid json\n" + err.Error())
	}

	return data
}

func (this *Config) populateConfigWithExtractedData(data map[string]interface{}) {
	if value, exists := data["servers"]; exists {
		if serverURLs, ok := value.([]interface{}); ok && len(serverURLs) > 0 {
			if _, ok := serverURLs[0].(string); !ok {
				panic("Server URLs should be of type string in the JSON config file!")
			}

			this.serverURLs = make([]string, len(serverURLs))

			for i := 0; i < len(serverURLs); i++ {
				this.serverURLs[i] = serverURLs[i].(string)
			}
		}
	}

	fmt.Print("hi")
}
