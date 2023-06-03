package configparser

import (
	"fmt"
	"os"
)

type Config struct {
	serverURLs                    []string
	balancingAlgorithmName        string
	serverTimeoutSeconds          int
	failedHealthChecksTillTimeout int
	slowStart                     bool
	stickySession                 bool
}

func (this *Config) InitDefault() {
	this.serverURLs = make([]string, 0)
	this.balancingAlgorithmName = "round-robin"
	this.serverTimeoutSeconds = 30
	this.failedHealthChecksTillTimeout = 3
	this.slowStart = false
	this.stickySession = false
}

func (this *Config) InitFromFile(configFilePath string) {
	if configFilePath == "" {
		panic("Config file path can not be empty")
	}

	if _, err:= os.Stat(configFilePath); err != nil {
		panic(err.Error())
	}

	this.InitDefault()

}

func (this *Config) Print() {
	fmt.Println(this)
}
