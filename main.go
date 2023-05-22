package main

import (
	"flag"
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
	flag.StringVar(&configFile, "f", "", "specify configuration file")
	flag.BoolVar(&testConfig, "t", false, "test configuration and exit")
	g, err := config.Parse(configFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if testConfig {
		log.Println("config parse success")
		return
	}
	s, err := server.NewServer(g)

	s.Run()
	s.PrintBaseConfig()
	log.Println("server start, listening")
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
