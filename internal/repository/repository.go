package repository

import (
	"CarRentalService/internal/models"

	"github.com/jmoiron/sqlx"
)



type Repository struct {
	Autorization
	User
	Car
	Rental
	Reviews
}


type Autorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email string, password string)(models.User, error)
}
type User interface {
	Delete(userId int) error
	GetInfo(userId int) (models.UserInfo, error)
	UpdateData(userUd int, input models.UserUpdate) (error)
}

type AvailabilityChecker interface {
    IsCarAvailable(carId int) (bool, error)
	GetCarById(carId int)(models.GetCarInfo, error)
}

type Car interface{
	AvailabilityChecker
	AddCar(userId int, car models.Car)(int, error)
	GetAllCars()([]models.BriefCarInfo, error)
	GetCarById(carId int)(models.GetCarInfo, error)
	DeleteCar(userId int, carId int) error
}

type Rental interface {
	AvailabilityChecker
	StartRental(userId int, totalPrice float64, input models.StartRent) (int, error)
	EndRental(rentId int)(float64, error)
	GetRentalInfo(rentId int) (models.Rental, error)
	RentalHistory(userId int) ([]models.RentalHistory, error)
}

type RentalCheck interface{
	GetRentalInfo(rentId int) (models.Rental, error)
}


type Reviews interface{
	LeaveReview(input models.Reviews)(int, error)
	ReviewExists(rentalId int)(bool, error)
}

func NewRepository(db *sqlx.DB) *Repository {
    carRepo := NewCarPostgres(db)
	rentalRepo := NewRentalPostgres(db, carRepo)
    return &Repository{
        Autorization: NewAuthPostgres(db),
        User: NewUserPostgres(db),
        Car: carRepo,
        Rental: NewRentalPostgres(db, carRepo),
		Reviews: NewReviewPostgres(db, rentalRepo),
    }
}