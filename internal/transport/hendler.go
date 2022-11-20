package transport

import (
	"encoding/json"
	"github.com/bolanebyla/Go-ToDo/internal/services"
	"log"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := services.GetTasks()

	err := Respond(w, tasks)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDto services.CreateTaskDto
	err := json.NewDecoder(r.Body).Decode(&taskDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		Respond(w, "Can't parse json")
		return
	}

	newTask := services.CreateTask(taskDto)

	Respond(w, newTask)
}
