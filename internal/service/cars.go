package service

import (
	"CarRentalService/internal/models"
	"CarRentalService/internal/repository"
	"fmt"
)


type CarService struct {
	repo repository.Car
}

func NewCarService(repo repository.Car) *CarService {
	return &CarService{repo: repo}
}

func (s *CarService) AddCar(userId int, car models.Car)(int, error) {
	return s.repo.AddCar(userId, car)
}

func (s *CarService) GetAllCars()([]models.BriefCarInfo, error){
	return s.repo.GetAllCars()
}

func (s *CarService) GetCarById(carId int)(models.GetCarInfo, error) {
	avaible, err := s.repo.IsCarAvailable(carId)
	if err != nil {
		return models.GetCarInfo{}, err
	}
	if !avaible {
		return models.GetCarInfo{}, fmt.Errorf("car is not avaible")
	}
	return s.repo.GetCarById(carId)
}

func (s *CarService) DeleteCar(userId, carId int) error {
	avaible, err := s.repo.IsCarAvailable(carId)
	if err != nil {
		return err
	}
	if !avaible {
		return fmt.Errorf("car is not avaible")
	}
	return s.repo.DeleteCar(userId, carId)
}