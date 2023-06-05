package configparser

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

const configFileName string = "test.config.json"

func createTemporaryTestFile(text string) error {
	file, err := os.Create(configFileName)

	if err != nil {
		return errors.New("Failed to create temporary test file")
	}

	_, err = file.Write([]byte(text))

	if err != nil {
		return errors.New("Failed to write to temporary test file")
	}

	err = file.Close()

	if err != nil {
		return errors.New("Failed to close temporary test file")
	}

	return nil
}

func panicCalled() bool {
	if r := recover(); r != nil {
		return true
	}

	return false
}

func TestNewConfig_FuncResultMatchesDefaultConfigConst(t *testing.T) {
	inConfig := NewConfig()
	wantConfig := Config{
		serverURLs:                    nil,
		balancingAlgorithmName:        "round_robin",
		serverTimeoutSeconds:          60,
		failedHealthChecksTillTimeout: 3,
		slowStart:                     false,
		slowStartSeconds:              120,
		stickySession:                 false,
	}

	if !reflect.DeepEqual(inConfig, wantConfig) {
		t.Errorf("TestNewConfig() == \n%+v\nwant \n%+v\n", inConfig, wantConfig)
	}
}

func TestNewConfig_FuncResultDifferentThanCustomConfig(t *testing.T) {
	inConfig := NewConfig()
	differentConfig := Config{
		serverURLs:                    make([]string, 0),
		balancingAlgorithmName:        "round_robin",
		serverTimeoutSeconds:          60,
		failedHealthChecksTillTimeout: 3,
		slowStart:                     false,
		slowStartSeconds:              120,
		stickySession:                 false,
	}

	if reflect.DeepEqual(inConfig, differentConfig) {
		t.Errorf("TestNewConfig() == \n%+v\nEXPECTED NOT EQUAL WITH\n%+v\n", inConfig, differentConfig)
	}

	differentConfig.serverURLs = nil
	differentConfig.balancingAlgorithmName = "least_connections"

	if reflect.DeepEqual(inConfig, differentConfig) {
		t.Errorf("TestNewConfig() == \n%+v\nEXPECTED NOT EQUAL WITH\n%+v\n", inConfig, differentConfig)
	}
}

func TestNewConfigFromFile_DefaultConfigProcessedFromFileSameAsDefault(t *testing.T) {
	err := createTemporaryTestFile(defaultConfig)

	if err != nil {
		t.Error(err.Error())
	}

	inConfig, err := NewConfigFromFile(configFileName)

	if err != nil {
		t.Error("Malformed default configuration read from file.\n" + err.Error())
	}

	wantConfig := Config{
		serverURLs:                    nil,
		balancingAlgorithmName:        "round_robin",
		serverTimeoutSeconds:          60,
		failedHealthChecksTillTimeout: 3,
		slowStart:                     false,
		slowStartSeconds:              120,
		stickySession:                 false,
	}

	if !reflect.DeepEqual(inConfig, wantConfig) {
		t.Errorf("TestNewConfig() == \n%+v\nwant \n%+v\n", inConfig, wantConfig)
	}

	err = os.Remove(configFileName)

	if err != nil {
		t.Error("Failed to remove temporary test file")
	}
}

func TestNewConfigFromFile_CustomConfigProcessedFromFileDifferentThanDefault(t *testing.T) {
	const customConfig string = `{
		"serverURLs": [
			"backend1.url.com",
			"backedn2.url.com"
		],
		"balancingAlgorithmName": "w_round_robit",
		"serverTimeoutSeconds": 120,
		"failedHealthChecksTillTimeout": 3,
		"slowStart": false,
		"slowStartSeconds": 120,
		"stickySession": true
	}`

	err := createTemporaryTestFile(customConfig)

	if err != nil {
		t.Error(err.Error())
	}

	inConfig, err := NewConfigFromFile(configFileName)

	if err != nil {
		t.Error("Malformed default configuration read from file.\n" + err.Error())
	}

	differentConfig := Config{
		serverURLs:                    nil,
		balancingAlgorithmName:        "round_robin",
		serverTimeoutSeconds:          60,
		failedHealthChecksTillTimeout: 3,
		slowStart:                     false,
		slowStartSeconds:              120,
		stickySession:                 false,
	}

	if reflect.DeepEqual(inConfig, differentConfig) {
		t.Errorf("TestNewConfig() == \n%+v\nEXPECTED NOT EQUAL WITH\n%+v\n", inConfig, differentConfig)
	}

	err = os.Remove(configFileName)

	if err != nil {
		t.Error("Failed to remove temporary test file")
	}
}

func TestNewConfigFromFile_ReturnNonFatalErrorOnMalformedConfigFile(t *testing.T) {
	const customConfig string = `{
		"serverURLs": [
			"backend1.url.com",
			"backedn2.url.com"
		],
		"balancingAlgorithmName": "w_round_robit",
		"serverTimeoutSeconds": 120,
		"failedHealthChecksTillTimeout": 3,
		"slowStart": false,
		"slowStartSeconds": 120,
		"unknownField": true
	}`

	err := createTemporaryTestFile(customConfig)

	if err != nil {
		t.Error(err.Error())
	}

	_, err = NewConfigFromFile(configFileName)

	if err == nil {
		t.Errorf("Not nil error expected when processing config file with an unknown field")
	}

	err = os.Remove(configFileName)

	if err != nil {
		t.Error("Failed to remove temporary test file")
	}
}

func TestNewConfigFromFile_ErrorOnConfigFileFieldWithMismatchedValue(t *testing.T) {	
	defer func() {
		if r := recover(); r == nil {
			t.Error("Panic expected due to malformed json config file field \"serverTimeoutSeconds\"")
		}
	}()

	defer func() {
		err := os.Remove(configFileName)

			if err != nil {
				t.Error("Failed to remove temporary test file")
			}
	}()
	
	const customConfig string = `{
		"serverURLs": [
			"backend1.url.com",
			"backedn2.url.com"
		],
		"balancingAlgorithmName": "w_round_robit",
		"serverTimeoutSeconds": true,
		"failedHealthChecksTillTimeout": 3,
		"slowStart": false,
		"slowStartSeconds": 120
	}`

	err := createTemporaryTestFile(customConfig)

	if err != nil {
		t.Error(err.Error())
	}

	_, err = NewConfigFromFile(configFileName)
}
