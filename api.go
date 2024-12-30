//server and apis
//use github.com/gorilla/mux for http router

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

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

	router.HandleFunc("/account", s.handleAccount)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
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
