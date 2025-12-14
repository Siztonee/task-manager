package services

import (
	"errors"
	"task-manager/internal/models"
	"task-manager/internal/repository"
	"task-manager/internal/utils"
)

type AuthService struct {
	Repo      *repository.UserRepository
	jwtSecret []byte
}

func NewAuthService(repo *repository.UserRepository, jwtSecret []byte) *AuthService {
	return &AuthService{
		Repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) GetAll() ([]models.User, error) {
	return s.Repo.GetAll()
}

func (s *AuthService) Register(user *models.User) error {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash

	return s.Repo.Create(user)
}

func (s *AuthService) Login(user *models.User) (string, error) {
	findedUser, err := s.Repo.FindByUsername(user.Username)
	if err != nil {
		return "", errors.New("invalid credentails")
	}

	if !utils.CheckPassword(findedUser.Password, user.Password) {
		return "", errors.New("invalid credentails")
	}

	return utils.GenerateToken(findedUser.ID, s.jwtSecret)

}
