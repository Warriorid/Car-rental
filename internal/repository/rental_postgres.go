package repository

import (
	"CarRentalService/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RentalPostgres struct {
	db *sqlx.DB
	carRepo AvailabilityChecker
}
var (
	ErrRentalNotFound = errors.New("rental not found")
	ErrRentalCompleted = errors.New("rental alredy completed")
)

func NewRentalPostgres(db *sqlx.DB, carRepo AvailabilityChecker) *RentalPostgres{
	return &RentalPostgres{db: db, carRepo: carRepo}
}

func (r *RentalPostgres) IsCarAvailable(carId int) (bool, error) {
    return r.carRepo.IsCarAvailable(carId)
}
func (r *RentalPostgres) GetCarById(carId int) (models.GetCarInfo, error) {
	return r.carRepo.GetCarById(carId)
}

func (r *RentalPostgres) StartRental(userId int, totalPrice float64, input models.StartRent) (int, error) {
	var rentId int

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	queryCar := "UPDATE cars SET is_available=false WHERE id=$1"
	if _, err := tx.Exec(queryCar, input.CarId); err!= nil {
		return 0, fmt.Errorf("error car available update: %s", err)
	}
	setValue := "(car_id, user_id, start_date, end_date, total_price, status)"
	queryRental := fmt.Sprintf("INSERT INTO rentals %s VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", setValue)
	row := tx.QueryRow(queryRental, input.CarId, userId, input.StartDate, input.EndDate, totalPrice, "active")
	if err := row.Scan(&rentId); err != nil {
		return 0, fmt.Errorf("error start rental, %s", err)
	}
	return rentId, tx.Commit()
}

func (r *RentalPostgres) EndRental(rentId int)(float64, error){
	tx, err := r.db.Begin()
	if err != nil {
		return 0.0, err
	}
	defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
            if err != nil {
                tx.Rollback()
            }
        }
    }()

	var totalPrice float64
	rental, err := r.GetRentalInfo(rentId)
	if err != nil {
		return 0.0, err
	}
	
	carQuery := "UPDATE cars SET is_available=true WHERE id=$1"
	if _, err := tx.Exec(carQuery, rental.CarId); err != nil {
		return 0.0, fmt.Errorf("error of update car status: %s", err)
	}
	rentStatusQuery := "UPDATE rentals SET status=$1 WHERE id=$2 AND status!='completed'"
	result, err := tx.Exec(rentStatusQuery, "completed", rentId)
	if err != nil {
		return 0.0, fmt.Errorf("error of updating rent status: %s", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0.0, err
	}
	if rowsAffected == 0 {
		return 0.0, ErrRentalCompleted
	}
	totalPrice = rental.TotalPrice


	return totalPrice, tx.Commit()
}

func (r *RentalPostgres) GetRentalInfo(rentId int) (models.Rental, error) {
	var rental models.Rental
	getRentQuery := "SELECT car_id, user_id, start_date, end_date, total_price, status FROM rentals WHERE id=$1"
	row := r.db.QueryRow(getRentQuery, rentId)
	if err := row.Scan(
		&rental.CarId, 
		&rental.UserId,
		&rental.StartDate,
		&rental.EndDate,
		&rental.TotalPrice,
		&rental.Status,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows){
				return rental, ErrRentalNotFound
			}
		return rental, fmt.Errorf("error of scan rental: %s", err)
	}
	return rental, nil
}

func(r *RentalPostgres) RentalHistory(userId int) ([]models.RentalHistory, error){
	query := `SELECT c.model as car, u.name as user, r.start_date, r.end_date, r.total_price, r.status
				from rentals r
				JOIN cars c on r.car_id = c.id
				join users u on r.user_id = u.id
				WHERE r.user_id = $1;`
	var rentalsList []models.RentalHistory
	if err := r.db.Select(&rentalsList, query, userId); err != nil {
		return rentalsList, fmt.Errorf("error of select from db: %s", err)
	}
	return rentalsList, nil
}
