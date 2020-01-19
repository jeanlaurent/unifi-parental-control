package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	api *API
	db  *kidsDB
}

func (s *server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func (s *server) ClientsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	// we should proxy probably here to the proper server
	clients, err := s.api.ListClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err) // should return 500s
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if ID == "" {
		json.NewEncoder(w).Encode(clients)
	} else {
		for index := 0; index < len(clients); index++ {
			if clients[index].ID == ID {
				json.NewEncoder(w).Encode(clients[index])
				break
			}
		}
	}
}

func (s *server) BlockedHandler(w http.ResponseWriter, r *http.Request) {
	blockedDevices, err := s.db.allBlocked()
	if err != nil {
		log.Fatalf("Fetching blocked: %v", err) // should return 500s
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blockedDevices)
}

func (s *server) BlockClientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	clients, err := s.api.ListClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err) // should return 500s
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for index := 0; index < len(clients); index++ {
		if clients[index].ID == ID {
			s.api.BlockClient("default", clients[index].MAC)
			json.NewEncoder(w).Encode(clients[index])
			s.db.addBlockedDevice(clients[index].Name, clients[index].MAC)
			break
		}
	}
}

func (s *server) UnblockClientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	clients, err := s.api.ListClients("default")
	if err != nil {
		log.Fatalf("Fetching clients: %v", err) // should return 500s
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for index := 0; index < len(clients); index++ {
		if clients[index].ID == ID {
			s.api.UnblockClient("default", clients[index].MAC)
			json.NewEncoder(w).Encode(clients[index])
			// s.db.deleteBlockedDevice(clients[index].Name, clients[index].MAC)
			break
		}
	}
}

func (s *server) start() {
	router := mux.NewRouter()
	router.HandleFunc("/clients/{ID}/unblock", s.UnblockClientHandler).Methods("PUT")
	router.HandleFunc("/clients/{ID}/block", s.BlockClientHandler).Methods("PUT")
	router.HandleFunc("/clients/{ID}", s.ClientsHandler).Methods("GET")
	router.HandleFunc("/clients/", s.ClientsHandler).Methods("GET")
	router.HandleFunc("/blocked/", s.BlockedHandler).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./app/")))
	log.Fatal(http.ListenAndServe(":8000", router))
}

func start(api *API, db *kidsDB) {
	server := server{api, db}
	server.start()
}
