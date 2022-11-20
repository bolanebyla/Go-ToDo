package services

import (
	"time"
)

func GetTasks() []GetTaskDto {
	mockTask1 := GetTaskDto{
		Id:        1,
		Text:      "Test 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockTask2 := GetTaskDto{
		Id:        2,
		Text:      "Test 2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockTasks := []GetTaskDto{mockTask1, mockTask2}

	return mockTasks
}

func CreateTask(taskDto CreateTaskDto) GetTaskDto {
	// save to db
	newTask := GetTaskDto{
		Id:        1234,
		Text:      taskDto.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return newTask
}
