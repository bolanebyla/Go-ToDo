package transport

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ApiServer struct {
	router mux.Router
	host   string
	port   string
}

func CreateApiServer(router *mux.Router, host string, port string) *ApiServer {
	return &ApiServer{
		router: *router,
		host:   host,
		port:   port,
	}
}

func (s *ApiServer) StartApiServer() {

	addr := s.host + ":" + s.port

	log.Printf("Start API server on: http://%s", addr)

	err := http.ListenAndServe(addr, &s.router)
	if err != nil {
		log.Fatal(err)
	}
}
