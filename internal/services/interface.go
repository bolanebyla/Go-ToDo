package services

type TaskRepo interface {
	CreateTask(dto *CreateTaskDto) *GetTaskDto
	GetTasks() *[]GetTaskDto
	GetTask(id int) *GetTaskDto
}
