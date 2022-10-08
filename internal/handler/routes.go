package handler

import (
	"net/http"

	"bitbucket.org/btcid/startrack/internal/service"
	"bitbucket.org/btcid/startrack/pkg/server"
)

func RegisterHandlers(s *server.Server, ctx *service.ServiceContext) {
	// Register Routes In Here
	routes := []server.Route{
		{
			Name:    "sample",
			Method:  http.MethodGet,
			Path:    "/sample",
			Handler: SampleHandler(ctx),
		},
		{
			Name:    "test",
			Method:  http.MethodGet,
			Path:    "/test",
			Handler: SampleHandler(ctx),
		},
	}

	s.AddRoutes(routes)
}
