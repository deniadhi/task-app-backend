package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID           uuid.UUID
  Name         string
	CreatedAt    time.Time
  UpdatedAt    time.Time
}

type Input struct {
	Name string
}