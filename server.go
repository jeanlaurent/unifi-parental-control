package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	api *API
}

func (s *server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func (s *server) ClientHandler(w http.ResponseWriter, r *http.Request) {
	// we should proxy probably here to the proper server
	clients, err := s.api.ListClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}

func (s *server) start() {
	r := mux.NewRouter()
	r.HandleFunc("/", s.HomeHandler)
	r.HandleFunc("/clients", s.ClientHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func start(api *API) {
	server := server{api}
	server.start()
}
