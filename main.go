package main

import (
	configparser "github.com/Alex-dev02/load_balancer/config_parser"
	loadbalancer "github.com/Alex-dev02/load_balancer/load_balancer"
)

func main() {
	loadbalancer.Hello()
	config := configparser.NewConfig()
	config.Print()
}
