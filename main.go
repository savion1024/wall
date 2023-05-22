package main

import (
	"flag"
	"fmt"
	"github.com/savion1024/wall/config"
	"github.com/savion1024/wall/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	testConfig bool
	configFile string
)

func main() {
	flag.StringVar(&configFile, "f", "", "Specify configuration file")
	flag.BoolVar(&testConfig, "t", false, "Test configuration and exit")
	g, err := config.Parse(configFile)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Config parse failed: %s", err.Error()))
	}
	if testConfig {
		log.Println("Config parse success")
		return
	}
	s := server.NewServer(g)

	s.Run()
	s.PrintBaseConfig()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
