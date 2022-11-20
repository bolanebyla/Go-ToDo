package transport

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func StartApiServer() {

	router := mux.NewRouter()

	router.HandleFunc("/api/tasks/", GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/", CreateTask).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	addr := host + ":" + port

	log.Printf("Start API server on: http://%s", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
	}
}
