package services

import (
	"errors"
	"task-manager/internal/http/dto"
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

func (s *TaskService) Update(id uint, req dto.UpdateTaskRequest) error {
	updates := map[string]interface{}{}
	if req.Title != nil {
		updates["title"] = req.Title
	}

	if req.Completed != nil {
		updates["completed"] = req.Completed
	}

	if len(updates) == 0 {
		return errors.New("Nothing to update")
	}

	return s.Repo.Update(id, updates)
}

func (s *TaskService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
