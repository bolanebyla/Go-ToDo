package services

import (
	"github.com/bolanebyla/Go-ToDo/internal/models"
	"time"
)

func GetTasks() []models.Task {
	mockTask1 := models.Task{
		Id:        1,
		Text:      "Test 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockTask2 := models.Task{
		Id:        2,
		Text:      "Test 2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockTasks := []models.Task{mockTask1, mockTask2}

	return mockTasks
}
