package transport

import (
	"github.com/bolanebyla/Go-ToDo/internal/services"
	"log"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := services.GetTasks()
	resp := Response()
	resp["data"] = tasks
	err := Respond(w, resp)
	if err != nil {
		log.Fatal(err)
	}
}
