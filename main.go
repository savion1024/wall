package main

import (
	"flag"
	"github.com/savion1024/wall/config"
	"log"
)

var (
	testConfig bool
	configFile string
)

func main() {
	flag.StringVar(&configFile, "f", "", "specify configuration file")
	flag.BoolVar(&testConfig, "t", false, "test configuration and exit")

	g, err := config.Parse(configFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
