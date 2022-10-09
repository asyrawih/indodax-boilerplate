package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"bitbucket.org/btcid/startrack/internal/config"
	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route

type Server struct {
	config *config.Config
}

// Use Gorilla Mux As Handler
var router *mux.Router

// Http
var server *http.Server

func MustNewServer(ctx *config.Config) *Server {
	return &Server{
		config: ctx,
	}
}

// AddRoutes method  
// Register Add Routes
func (s *Server) AddRoutes(routes Routes) {
	router = mux.NewRouter().StrictSlash(false)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}
}

// Route method  
// Just in Case need access intance of router for modified on Route
func (s *Server) Route() *mux.Router {
	return router
}

// Start Http Server
func (s *Server) Start() {
	log.Printf("server:[%s] has Started on port %s", s.config.Name, s.config.Port)

	server = &http.Server{
		Addr: s.config.Host + ":" + s.config.Port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}
