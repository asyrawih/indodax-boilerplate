package handler

import (
	"net/http"

	"bitbucket.org/btcid/startrack/internal/service"
)

// SampleHandler function  
func SampleHandler(svc *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(svc.Config.Name))
	}
}

// HelloWorld function  
func HelloWorld(svc *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(svc.Config.Name))
	}
}
