package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID           uuid.UUID
  Detail       string
  Assignee     string
	Deadline		 string
	Finished		 bool
	CreatedAt    time.Time
  UpdatedAt    time.Time
}

type Input struct {
	Detail       string
  Assignee     string
	Deadline		 string
	Finished		 bool
}