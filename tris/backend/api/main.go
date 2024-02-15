package main

import (
	"encoding/json"
	"net/http"

	//"github.com/gorilla/websocket"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

type GameState struct {
	XIsNext bool    `json:"xIsNext"`
	Squares []string `json:"squares"`
	Winner  string   `json:"winner"`
}

var gameState GameState
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan GameState)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
func main() {
	// Inizializza lo stato del gioco
	gameState = GameState{
		XIsNext: true,
		Squares: make([]string, 9),
		Winner:  "",
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/game-state", getGameStateHandler)
	mux.HandleFunc("/make-move", makeMoveHandler)
	mux.HandleFunc("/reset-game", resetGameHandler)
	mux.HandleFunc("/ws", handleWebSocket)

	handler := cors.Default().Handler(mux)
	go handleBroadcast()
	http.ListenAndServe(":8080", handler)

}
func getGameStateHandler(w http.ResponseWriter, r *http.Request) {
	// Restituisci lo stato corrente del gioco come JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameState)
}

func makeMoveHandler(w http.ResponseWriter, r *http.Request) {
	// Decodifica la mossa del giocatore dal corpo della richiesta JSON
	var moveData map[string][]string
	json.NewDecoder(r.Body).Decode(&moveData)

	// Aggiorna lo stato del gioco con la mossa del giocatore
	gameState.Squares = moveData["squares"]
	gameState.XIsNext = !gameState.XIsNext
	gameState.Winner = calculateWinner(gameState.Squares)

	// Restituisci il nuovo stato del gioco come JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameState)
}

func resetGameHandler(w http.ResponseWriter, r *http.Request) {
	// Reimposta lo stato del gioco a quello iniziale
	gameState.XIsNext = true
	gameState.Squares = make([]string, 9)
	gameState.Winner = ""

	// Restituisci il nuovo stato del gioco come JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameState)
}
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		// Read message from the client
		_, _, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			break
		}
	}
}
func handleBroadcast() {
	for {
		state := <-broadcast
		for client := range clients {
			err := client.WriteJSON(state)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
func calculateWinner(squares []string) string {
	lines := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for i := 0; i < len(lines); i++ {
		a, b, c := lines[i][0], lines[i][1], lines[i][2]
		if squares[a] != "" && squares[a] == squares[b] && squares[a] == squares[c] {
			return squares[a]
		}
	}
	return ""
}
