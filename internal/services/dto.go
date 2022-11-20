package services

import "time"

type CreateTaskDto struct {
	Text string `json:"text"`
}

type GetTaskDto struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
