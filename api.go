//server and apis
//use github.com/gorilla/mux for http router

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Convert to JSON
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// make our APIs work with HandlerFunc from Mux
type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle error
			writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})

		}
	}

}

// API Server
type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	log.Panicln("API Server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// APIs
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.getAccount(w, r)
	}

	if r.Method == "POST" {
		return s.createAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.deleteAccount(w, r)
	}

	return fmt.Errorf("method not supported %s", r.Method)
}

func (s *APIServer) getAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) createAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) deleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
