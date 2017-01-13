package responses

import (
	"encoding/json"
	"net/http"
	"time"
)

type Message struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func WriteMessage(message string, status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(Message{message, time.Now()})
}
