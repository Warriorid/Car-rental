package repository

import (
	"CarRentalService/internal/models"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CarPostgres struct {
	db *sqlx.DB
}


func NewCarPostgres(db *sqlx.DB) *CarPostgres {
	return &CarPostgres{db: db}
}


func (r *CarPostgres) AddCar(userId int, car models.Car) (int, error) {
	var carId int
	addCarQuery := fmt.Sprintf("INSERT INTO %s (model, year, color, mileage, price_per_day, is_available, location, owner_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", "cars")
	row := r.db.QueryRow(
		addCarQuery, 
		car.Model,
		car.Year,
		car.Color,
		car.Mileage,
		car.PricePerDay,
		car.IsAvailable,
		car.Location,
		userId)
	if err := row.Scan(&carId); err != nil {
		return 0, fmt.Errorf("failed to add car: %v", err)
	}
	
	return carId, nil
}

func (r *CarPostgres) GetAllCars()([]models.BriefCarInfo, error) {
	carInfo := make([]models.BriefCarInfo, 0)
	query := fmt.Sprintf("SELECT model, color, price_per_day FROM %s WHERE is_available=true", "cars")
	if err := r.db.Select(&carInfo, query); err != nil {
		return nil, fmt.Errorf("error of read from db:%d", err)
	}
	return carInfo, nil

}

func (r *CarPostgres) IsCarAvailable(carId int)(bool, error){
	var isAvaible bool

	query := fmt.Sprint("SELECT is_available FROM cars WHERE id=$1")
	row := r.db.QueryRow(query,carId)
	if err := row.Scan(&isAvaible); err != nil{
		return false, fmt.Errorf("failed to check car availability: %w", err)
	}
	return isAvaible, nil
}

func (r *CarPostgres) GetCarById(carId int)(models.GetCarInfo, error) {
	var car models.GetCarInfo
	setQuery := "c.model, c.year, c.color, c.mileage, c.price_per_day, c.location, u.name"
	query := fmt.Sprintf("SELECT %s from %s c join %s u on u.id = c.owner_id where c.id = $1",
	 						setQuery, "cars", "users")
	err := r.db.QueryRow(query, carId).Scan(
        &car.Model,
        &car.Year,
        &car.Color,
        &car.Mileage,
        &car.PricePerDay,
        &car.Location,
        &car.OwnerName,
    )
	if err != nil {
		if err == sql.ErrNoRows {
            return car, sql.ErrNoRows
        }
		return car, fmt.Errorf("error of scan: %s", err)
	}
	return car, nil
}

func (r *CarPostgres) DeleteCar(userId, carId int) error {

	query := fmt.Sprint("DELETE FROM cars c USING users u WHERE u.id = $1 and u.id = c.owner_id AND c.id = $2")
	result, err := r.db.Exec(query, userId, carId)
	rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("err in sql query:%s", err)
    }

    if rowsAffected == 0 {
        return sql.ErrNoRows
    }

	return err
	
}