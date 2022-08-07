package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FetchAll() ([]Task, error)
	Save(input Input) (Task, error)
	FetchByID(ID string) (Task, error)
	Update(id string, payload Input) (Task, error)
	Delete(id string) (Task, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FetchAll() ([]Task, error) {
	var tasks []Task

	err := r.db.Order("created_at desc").Find(&tasks).Error

	return tasks, err
}

func (r *repository) FetchByID(ID string) (Task, error) {
	convertedID, _ := uuid.Parse(ID)
	var task Task

	err := r.db.Find(&task, convertedID).Error

	return task, err
}

func (r *repository) Save(input Input) (Task, error) {
	var task Task
	task.Detail = input.Detail
	task.Assignee = input.Assignee
	task.Deadline = input.Deadline
	task.Finished = input.Finished
	task.ID = uuid.New()
	err := r.db.Create(&task).Error

	return task, err
}

func (r *repository) Update(id string, payload Input) (Task, error) {
	convertedID, _ := uuid.Parse(id)
	var task Task

	r.db.Find(&task, convertedID)

	task.Detail = payload.Detail
	task.Assignee = payload.Assignee
	task.Deadline = payload.Deadline
	task.Finished = payload.Finished
	err := r.db.Save(&task).Error

	return task, err
}

func (r *repository) Delete(id string) (Task, error) {
	convertedID, _ := uuid.Parse(id)
	var task Task

	r.db.Find(&task, convertedID)
	err := r.db.Delete(&task).Error

	return task, err
}
