package service

import (
	"CarRentalService/internal/models"
	"CarRentalService/internal/repository"
	"errors"
	"fmt"
)

var ErrPermission = errors.New("You have not permission complete rent")

type RentalService struct{
	carRepo   repository.AvailabilityChecker
	repo repository.Rental
}

func NewRentalServise(carRepo repository.AvailabilityChecker, repo repository.Rental) *RentalService{
	return &RentalService{carRepo: carRepo, repo: repo}
}

func (s *RentalService) StartRental(userId int, input models.StartRent) (int, error) {
	avaible, err := s.repo.IsCarAvailable(input.CarId)
	if err != nil {
		return 0, err
	}
	if !avaible {
		return 0, fmt.Errorf("car is not avaible")
	}
	
	if err := checkDate(input); err != nil {
		return 0, err
	}
	
	car, err := s.carRepo.GetCarById(input.CarId)
	if err != nil {
		return 0, err
	}
	totalPrice, err := priceCalculation(input.StartDate, input.EndDate, car.PricePerDay)
	if err != nil {
		return 0.0, err
	}
	return s.repo.StartRental(userId, totalPrice, input)
}

func (s *RentalService) EndRental(rentId, userId int) (float64, error) {
	rental, err := s.repo.GetRentalInfo(rentId)
	if err != nil {
		return 0.0, err
	}
	if rental.UserId != userId {
		return 0.0, ErrPermission
	}
	return s.repo.EndRental(rentId)
}

func (s *RentalService) RentalHistory(userId int)([]models.RentalHistory, error){
	return s.repo.RentalHistory(userId)
}