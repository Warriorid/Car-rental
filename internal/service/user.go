package service

import (
	"CarRentalService/internal/models"
	"CarRentalService/internal/repository"
)


type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s * UserService) GetInfo(userId int) (models.UserInfo, error) {
	return s.repo.GetInfo(userId)
}

func (s *UserService) UpdateData(userId int, input models.UserUpdate) error {
	return s.repo.UpdateData(userId, input)
}