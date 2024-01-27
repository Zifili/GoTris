package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Activities *Activities
}

type IDDocument struct {
	ID uint64 `json:"id"`
}

type ActivityDocument struct {
	Activity Activity `json:"activity"`
}

func (s *httpServer) handlePost(w http.ResponseWriter, r *http.Request) {
	var req ActivityDocument
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := s.Activities.Insert(req.Activity)
	res := IDDocument{ID: id}
	json.NewEncoder(w).Encode(res)
}
func (s *httpServer) handleGet(w http.ResponseWriter, r *http.Request) {
	var req IDDocument
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	activity, err := s.Activities.Retrieve(req.ID)
	if err == ErrIDNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := ActivityDocument{Activity: activity}
	json.NewEncoder(w).Encode(res)
}
func NewHTTPServer(addr string) *http.Server {

	server := &httpServer{
		Activities: &Activities{},
	}
	r := mux.NewRouter()
	r.HandleFunc("/", server.handlePost).Methods("POST")
	r.HandleFunc("/", server.handleGet).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
