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

func (r *TaskRepository) Update(id uint, updates map[string]interface{}) error {
	res := r.DB.
		Model(&models.Task{}).
		Where("id = ?", id).
		Updates(updates)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *TaskRepository) Delete(id uint) error {
	res := r.DB.Delete(&models.Task{}, id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
