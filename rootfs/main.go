package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PluginActivateResponse struct {
	Implements []string `json:"Implements"`
}

func PluginActivate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PluginActivate called")

	response := PluginActivateResponse{
		Implements: []string{"NetworkDriver"},
	}

	json.NewEncoder(w).Encode(response)
}

type CreateNetworkRequest struct {
	NetworkID string `json:"NetworkID"`
}

func CreateNetwork(w http.ResponseWriter, r *http.Request) {
	var req CreateNetworkRequest
	json.NewDecoder(r.Body).Decode(&req)

	fmt.Printf("CreateNetwork called: %v\n", req)

	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/Plugin.Activate", PluginActivate).Methods("POST")
	r.HandleFunc("/NetworkDriver.CreateNetwork", CreateNetwork).Methods("POST")

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

