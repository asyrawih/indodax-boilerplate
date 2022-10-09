package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct{}

// ToJson function  
// Convert Slice Into Json
func ToJson(r http.ResponseWriter, data interface{}) {
	r.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(data)

	if err != nil {
		log.Fatalf("%s", err)
	}
	_, err2 := r.Write(b)

	if err != nil {
		log.Fatalf("%s", err2)
	}
}
