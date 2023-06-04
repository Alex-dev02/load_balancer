package configparser

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
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
