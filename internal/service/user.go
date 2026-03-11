package service

import (
	"main/internal/models"
	"main/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService{
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll() ([]models.User, error){
	users, err := s.repo.GetUser()
	if err != nil{
		return  nil, err
	}
	
	return users, nil
}

