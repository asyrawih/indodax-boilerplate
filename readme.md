# Start Track

- [x] Http Api
- [x] Register Routes
- [ ] Register Middleware
- [ ] Implements Grpc API And Running Both


### How To Start Http Server 

```go 

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

var port = flag.String("p", "3000", "set for running port")

func main() {
	flag.Parse()

	var config config.Config

	if err := conf.LoadConfig(*configFile, &config); err != nil {
		log.Fatal(err)
	}

	// Need Set The Optional Port Affter the Config Successfully Loaded
	config.Port = *port

	// Spawn Server
	srv := server.MustNewServer(&config)

	// Register Service In Here
	ctx := service.NewServiceContext(&config)
	handler.RegisterHandlers(srv, ctx)

	go srv.Start()

	// Can Add Grpc Server In Here
	// And Spawn int Another Goroutine
	// Block the main Goroutine
	select {}

}

```

```bash
# Option
  -f string
        the config file (default "etc/startrack.yaml")
  -p string
        set for running port (default "3000")

```
