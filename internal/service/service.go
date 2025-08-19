package service

import (
	"CarRentalService/internal/models"
	"CarRentalService/internal/repository"
)


type Service struct {
	Autorization
	User
	Car
	Rental
	Reviews
}

type Autorization interface {
	CreateUser(user models.User)(int, error)
	GenerateToken(email string, password string)(string, error)
	ParseToken(tokenString string)(int, error)
}
type User interface {
	Delete(userId int) error
	GetInfo(userId int) (models.UserInfo, error)
	UpdateData(userUd int, input models.UserUpdate) (error)
}
type Car interface {
	AddCar(userId int, car models.Car)(int, error)
	GetAllCars()([]models.BriefCarInfo, error)
	GetCarById(carId int)(models.GetCarInfo, error)
	DeleteCar(userId int, carId int) error
}
type Rental interface{
	StartRental(userId int, input models.StartRent) (int, error )
	EndRental(rentId, userId int)(float64, error)
	RentalHistory(userId int) ([]models.RentalHistory, error)
}
type Reviews interface{
	LeaveReview(userId int, input models.Reviews) (int, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		User: NewUserService(repos.User),
		Car: NewCarService(repos.Car),
		Rental: NewRentalServise(repos.Car, repos.Rental),
		Reviews: NewReviewServise(repos.Reviews, repos.Rental),
	}
}
