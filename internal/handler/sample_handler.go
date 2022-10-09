package handler

import (
	"net/http"

	"bitbucket.org/btcid/startrack/internal/service"
	"bitbucket.org/btcid/startrack/pkg/response"
)

// SampleHandler function  
func SampleHandler(svc *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hello := map[string]string{"hello": "hello"}
		response.ToJson(w, hello)
	}
}

// HelloWorld function  
func HelloWorld(svc *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		hello := map[string]string{"hello": "hello"}
		response.ToJson(w, hello)
	}
}
