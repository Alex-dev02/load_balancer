package main

import (
	"github.com/Alex-dev02/load_balancer/load_balancer"
	"github.com/Alex-dev02/load_balancer/config_parser"
)

func main() {
	loadbalancer.Hello()
	config := configparser.NewConfigFromFile("example.config.json")
	config.Print()
}
