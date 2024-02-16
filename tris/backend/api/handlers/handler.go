package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/websocket"
)

type GameState struct {
	XIsNext bool     `json:"xIsNext"`
	Squares []string `json:"squares"`
	Winner  string   `json:"winner"`
}

var Gamestate GameState
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan GameState)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetGameStateHandler(w http.ResponseWriter, r *http.Request) {
	// Restituisci lo stato corrente del gioco come JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Gamestate)
}
func MakeMoveHandler(w http.ResponseWriter, r *http.Request) {
	// Decodifica la mossa del giocatore dal corpo della richiesta JSON
	var moveData map[string][]string
	json.NewDecoder(r.Body).Decode(&moveData)

	// Aggiorna lo stato del gioco con la mossa del giocatore
	Gamestate.Squares = moveData["squares"]
	Gamestate.XIsNext = !Gamestate.XIsNext
	Gamestate.Winner = calculateWinner(Gamestate.Squares)

	// Restituisci il nuovo stato del gioco come JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Gamestate)
}
func ResetGameHandler(w http.ResponseWriter, r *http.Request) {
	// Reimposta lo stato del gioco a quello iniziale
	Gamestate.XIsNext = true
	Gamestate.Squares = make([]string, 9)
	Gamestate.Winner = ""

	// Restituisci il nuovo stato del gioco come JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Gamestate)
}
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
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
func HandleBroadcast() {
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
