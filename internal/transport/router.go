package transport

import "github.com/gorilla/mux"

func CreateRouter(taskHandler *TaskHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/tasks/", taskHandler.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id:[0-9]+}/", taskHandler.GetTask).Methods("GET")
	router.HandleFunc("/api/tasks/", taskHandler.CreateTask).Methods("POST")

	return router
}
