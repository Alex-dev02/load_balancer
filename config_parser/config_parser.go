package configparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const defaultConfig string = `{
	"serverURLs": [],
	"balancingAlgorithmName": "round_robin",
	"serverTimeoutSeconds": 60,
	"failedHealthChecksTillTimeout": 3,
	"slowStart": false,
	"slowStartSeconds": 120,
	"stickySession": false
}`

type Config struct {
	serverURLs                    []string
	balancingAlgorithmName        string
	serverTimeoutSeconds          int
	failedHealthChecksTillTimeout int
	slowStart                     bool
	slowStartSeconds              int
	stickySession                 bool
}

func NewConfig() Config {
	var c Config

	data := unmarshalData([]byte(defaultConfig))
	c.populateConfigWithExtractedData(data)

	return c
}

func NewConfigFromFile(configFilePath string) Config {
	if configFilePath == "" {
		panic("Config file path can not be empty")
	}

	if _, err := os.Stat(configFilePath); err != nil {
		panic(err.Error())
	}

	c := NewConfig()
	c.loadConfigFromFile(configFilePath)

	return c
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
	if value, exists := data["serverURLs"]; exists {
		if serverURLs, ok := value.([]interface{}); ok && len(serverURLs) > 0 {
			if _, ok := serverURLs[0].(string); !ok {
				panic("serverURLs must be of type string in the JSON config file!")
			}

			this.serverURLs = make([]string, len(serverURLs))

			for i := 0; i < len(serverURLs); i++ {
				this.serverURLs[i] = serverURLs[i].(string)
			}
		}
	}

	if value, exists := data["balancingAlgorithmName"]; exists {
		if algorithm, ok := value.(string); ok {
			this.balancingAlgorithmName = algorithm
		} else {
			panic("balancingAlgorithmName must be of type string in the JSON config file!")
		}
	}

	if value, exists := data["serverTimeoutSeconds"]; exists {
		if timeoutSeconds, ok := value.(float64); ok {
			this.serverTimeoutSeconds = int(timeoutSeconds)
		} else {
			panic("serverTimeoutSeconds must be of type int in the JSON config file!")
		}
	}

	if value, exists := data["failedHealthChecksTillTimeout"]; exists {
		if failedHealthCheckAttempts, ok := value.(float64); ok {
			this.failedHealthChecksTillTimeout = int(failedHealthCheckAttempts)
		} else {
			panic("failedHealthChecksTillTimeout must be of type int in the JSON config file!")
		}
	}

	if value, exists := data["slowStart"]; exists {
		if slowStart, ok := value.(bool); ok {
			this.slowStart = slowStart
		} else {
			panic("slowStart must be of type bool in the JSON config file!")
		}
	}

	if value, exists := data["slowStartSeconds"]; exists {
		if slowStartSeconds, ok := value.(float64); ok {
			this.slowStartSeconds = int(slowStartSeconds)
		} else {
			panic("slowStartSeconds must be of type int in the JSON config file!")
		}
	}

	if value, exists := data["stickySession"]; exists {
		if stickySession, ok := value.(bool); ok {
			this.stickySession = stickySession
		} else {
			panic("stickySession must be of type bool in the JSON config file!")
		}
	}
}
