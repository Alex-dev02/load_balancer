package configparser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

// Config struct encapsulates all the needed configuration data in order for a load balancer
// to work properly. It is also used for enabling/disabling certain non-vital features
type Config struct {
	serverURLs                    []string
	balancingAlgorithmName        string
	serverTimeoutSeconds          uint
	failedHealthChecksTillTimeout uint
	slowStart                     bool
	slowStartSeconds              uint
	stickySession                 bool
}

// NewConfig initializes the Config struct with a predefined default configuration.
// It returns the initialized Config struct
func NewConfig() Config {
	var c Config

	data := unmarshalData([]byte(defaultConfig))
	c.populateConfigWithExtractedData(data)

	return c
}

// NewConfigFromFile initializes the Config struct with a custom configuration
// prvided in the specified file.
// It returns the initialized Config struct and an error in case the data validation went wrong
func NewConfigFromFile(configFilePath string) (Config, error) {
	if configFilePath == "" {
		panic("Config file path can not be empty")
	}

	if _, err := os.Stat(configFilePath); err != nil {
		panic(err.Error())
	}

	c := NewConfig()
	err := c.loadConfigFromFile(configFilePath)

	return c, err
}

func (this *Config) Print() {
	fmt.Println(this)
}

func (this *Config) loadConfigFromFile(configFilePath string) error {
	configDataJson := extractDataFromFile(configFilePath)
	data := unmarshalData(configDataJson)
	err := this.populateConfigWithExtractedData(data)

	return err
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

func (this *Config) populateConfigWithExtractedData(data map[string]interface{}) error {
	if value, exists := data["serverURLs"]; exists {
		delete(data, "serverURLs")

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
		delete(data, "balancingAlgorithmName")

		if algorithm, ok := value.(string); ok {
			this.balancingAlgorithmName = algorithm
		} else {
			panic("balancingAlgorithmName must be of type string in the JSON config file!")
		}
	}

	if value, exists := data["serverTimeoutSeconds"]; exists {
		delete(data, "serverTimeoutSeconds")

		if timeoutSeconds, ok := value.(float64); ok && timeoutSeconds >= 0{
			this.serverTimeoutSeconds = uint(timeoutSeconds)
		} else {
			panic("serverTimeoutSeconds must be of type uint in the JSON config file!")
		}
	}

	if value, exists := data["failedHealthChecksTillTimeout"]; exists {
		delete(data, "failedHealthChecksTillTimeout")

		if failedHealthCheckAttempts, ok := value.(float64); ok && failedHealthCheckAttempts >= 0 {
			this.failedHealthChecksTillTimeout = uint(failedHealthCheckAttempts)
		} else {
			panic("failedHealthChecksTillTimeout must be of type uint in the JSON config file!")
		}
	}

	if value, exists := data["slowStart"]; exists {
		delete(data, "slowStart")

		if slowStart, ok := value.(bool); ok {
			this.slowStart = slowStart
		} else {
			panic("slowStart must be of type bool in the JSON config file!")
		}
	}

	if value, exists := data["slowStartSeconds"]; exists {
		delete(data, "slowStartSeconds")

		if slowStartSeconds, ok := value.(float64); ok && slowStartSeconds >= 0 {
			this.slowStartSeconds = uint(slowStartSeconds)
		} else {
			panic("slowStartSeconds must be of type uint in the JSON config file!")
		}
	}

	if value, exists := data["stickySession"]; exists {
		delete(data, "stickySession")

		if stickySession, ok := value.(bool); ok {
			this.stickySession = stickySession
		} else {
			panic("stickySession must be of type bool in the JSON config file!")
		}
	}

	if len(data) > 0 {
		var unknownKeys strings.Builder
		for key := range data {
			_, err := unknownKeys.WriteString(key)
			_, err2 := unknownKeys.WriteString(" ")

			if err != nil || err2 != nil {
				panic("Coulnd not write to Builder while processing unknown json fields from config file")
			}
		}
		return errors.New("Non-fatal: Unknown fields found in json config file:\n" + unknownKeys.String())
	}

	return nil
}
