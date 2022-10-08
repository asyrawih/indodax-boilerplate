package main

import (
	"flag"
	"log"

	"bitbucket.org/btcid/startrack/internal/config"
	"bitbucket.org/btcid/startrack/internal/handler"
	"bitbucket.org/btcid/startrack/internal/service"
	"bitbucket.org/btcid/startrack/pkg/conf"
	"bitbucket.org/btcid/startrack/pkg/server"
)

var configFile = flag.String("f", "etc/startrack.yaml", "the config file")

func main() {
	flag.Parse()

	var config config.Config

	if err := conf.LoadConfig(*configFile, &config); err != nil {
		log.Fatal(err)
	}
	// Spawn Server
	srv := server.MustNewServer(&config)

	// Register Service In Here
	ctx := service.NewServiceContext(&config)
	handler.RegisterHandlers(srv, ctx)

	go srv.Start()

	// Can Add Grpc Server In Here
	// And Spawn int Another Goroutine
	select {}

}
