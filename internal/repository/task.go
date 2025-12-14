package repository

import (
	"task-manager/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.DB.Create(task).Error
}
