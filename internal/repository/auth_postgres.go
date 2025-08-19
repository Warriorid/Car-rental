package repository

import (
	"CarRentalService/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user models.User)(int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, phone, driver_license, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id", "users")
	
	row := r.db.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Phone,
		user.DriverLicense,
		user.Password,
	 )
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to create user %s", err.Error())
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email string, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id, name, email, phone, driver_license FROM %s WHERE email=$1 AND password_hash=$2", "users")
	err := r.db.Get(&user, query, email, password)
	return user, err
}