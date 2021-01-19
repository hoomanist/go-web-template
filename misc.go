package main

import (
	"encoding/json"
	"net/http"
)

func HandlePingRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(map[string]interface{}{
			"message": "pong",
		})
		w.Write([]byte(response))

	}
}
