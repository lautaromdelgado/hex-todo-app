package model

import "time"

type Task struct {
	ID        uint      `json:"id"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	CreatedIn time.Time `json:"createdIn"`
}
