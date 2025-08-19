package repository

import (
	"CarRentalService/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ReviewsPostgres struct {
	db *sqlx.DB
	rentalRepo RentalCheck
}

func NewReviewPostgres(db *sqlx.DB, rentalRepo RentalCheck) *ReviewsPostgres {
	return &ReviewsPostgres{db: db, rentalRepo: rentalRepo}
}

func (r *ReviewsPostgres) GetRentalInfo(rentId int) (models.Rental, error) {
	return r.rentalRepo.GetRentalInfo(rentId)
}

func (r *ReviewsPostgres) LeaveReview(input models.Reviews) (int, error){
	
	var reviewId int
	query := "INSERT INTO reviews (rental_id, rating, comment) VALUES ($1, $2, $3) RETURNING id"
	row := r.db.QueryRow(
		query,
		input.RentalId,
		input.Rating,
		input.Comment,
	)
	if err := row.Scan(&reviewId); err != nil {
		return 0, fmt.Errorf("error of scan reviewId: %s", err)
	}
	return reviewId, nil
}

func (r *ReviewsPostgres) ReviewExists(rentalId int)(bool, error){
	var reviewExists bool
	query := "SELECT * FROM reviews WHERE id=$1"
	result, err := r.db.Exec(query, rentalId)
	if err != nil {
		return reviewExists, fmt.Errorf("error review exists: %s", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return reviewExists, err
	}
	if rowsAffected == 0 {
		return false, nil
	}
	return true, nil
}