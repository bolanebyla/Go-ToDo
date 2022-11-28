package database

import (
	"github.com/bolanebyla/Go-ToDo/internal/services"
	"time"
)

type TaskRepo struct {
}

func CreateTaskRepo() *TaskRepo {
	return &TaskRepo{}
}

func (r TaskRepo) GetTasks() *[]services.GetTaskDto {
	mockTask1 := services.GetTaskDto{
		Id:        1,
		Text:      "Test 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockTask2 := services.GetTaskDto{
		Id:        2,
		Text:      "Test 2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockTasks := []services.GetTaskDto{mockTask1, mockTask2}

	return &mockTasks
}

func (r TaskRepo) CreateTask(dto *services.CreateTaskDto) *services.GetTaskDto {
	newTask := services.GetTaskDto{
		Id:        1234,
		Text:      dto.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &newTask
}

func (r TaskRepo) GetTask(id int) *services.GetTaskDto {
	task := services.GetTaskDto{
		Id:        id,
		Text:      "Test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &task
}
