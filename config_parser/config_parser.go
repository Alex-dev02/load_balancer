package configparser

import (
	"fmt"
)

type Config struct {
	serverURLs                    []string
	balancingAlgorithmName        string
	serverTimeoutSeconds          int
	failedHealthChecksTillTimeout int
	slowStart                     bool
	stickySession                 bool
}

func (this *Config) Init() {
	this.serverURLs = make([]string, 0)
	this.balancingAlgorithmName = "round-robin"
	this.serverTimeoutSeconds = 30
	this.failedHealthChecksTillTimeout = 3
	this.slowStart = false
	this.stickySession = false
}

func (this *Config) Print() {
	fmt.Println(this)
}
