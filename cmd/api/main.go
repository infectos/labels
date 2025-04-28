package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

const version = "1.0.0"

const (
	RoomCodeLength  = 4
	MaxRoomCapacity = 10
	RoomTTL         = 6 * time.Hour
	ShutdownTimeout = 10 * time.Second
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type GameServer struct {
	router *chi.Mux
	server *http.Server
	// rooms    map[string]*GameRoom
	roomsMux sync.RWMutex
	wg       sync.WaitGroup
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

type hello struct {
	Text string `json:"text"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	hello := &hello{"test"}
	respondWithJSON(w, 200, hello)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Exported Handler that Vercel requires
var Handler = http.HandlerFunc(greet)
