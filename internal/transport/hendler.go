package transport

import (
	"encoding/json"
	"github.com/bolanebyla/Go-ToDo/internal/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	service services.TaskService
}

func CreateTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: *service}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.service.GetTasks()

	err := Respond(w, tasks)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		Respond(w, "Переданный параметр id не является целым числом")
		return
	}

	task := h.service.GetTask(taskId)

	if task == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Respond(w, task)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDto services.CreateTaskDto
	err := json.NewDecoder(r.Body).Decode(&taskDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		Respond(w, "Can't parse json")
		return
	}

	newTask := h.service.CreateTask(&taskDto)

	Respond(w, newTask)
}
