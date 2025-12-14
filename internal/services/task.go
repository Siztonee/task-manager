package services

import (
	"task-manager/internal/models"
	"task-manager/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) GetAll() ([]models.Task, error) {
	return s.Repo.GetAll()
}

func (s *TaskService) Create(task *models.Task) error {
	return s.Repo.Create(task)
}
