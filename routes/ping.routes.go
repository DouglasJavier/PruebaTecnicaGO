package routes

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Message: "pong"}
	json.NewEncoder(w).Encode(response)

}
