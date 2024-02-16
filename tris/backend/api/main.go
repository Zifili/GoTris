package main

import (
	"gotris/api/handlers"
	"net/http"
	"github.com/rs/cors"
)

func main() {
	// Init game
	handlers.Gamestate = handlers.GameState{
		XIsNext: true,
		Squares: make([]string, 9),
		Winner:  "",
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/game-state", handlers.GetGameStateHandler)
	mux.HandleFunc("/make-move", handlers.MakeMoveHandler)
	mux.HandleFunc("/reset-game", handlers.ResetGameHandler)
	mux.HandleFunc("/ws", handlers.HandleWebSocket)

	handler := cors.Default().Handler(mux)
	go handlers.HandleBroadcast()
	http.ListenAndServe(":8080", handler)
}