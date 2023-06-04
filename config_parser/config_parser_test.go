package configparser

import (
	"os"
	"reflect"
	"testing"
)

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
	configFileName := "test.config.json"
	file, err := os.Create(configFileName)

	if err != nil {
		t.Error("Failed to create temporary test file")
	}

	_, err = file.Write([]byte(defaultConfig))

	if err != nil {
		t.Error("Failed to write to temporary test file")
	}

	err = file.Close()

	if err != nil {
		t.Error("Failed to close temporary test file")
	}

	inConfig := NewConfigFromFile(configFileName)
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
