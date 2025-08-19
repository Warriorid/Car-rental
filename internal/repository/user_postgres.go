package repository

import (
	"CarRentalService/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct{
	db *sqlx.DB
}
func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}


func (r * UserPostgres) Delete(userId int) error {
	
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", "users")
	result, err := r.db.Exec(query, userId)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r *UserPostgres) GetInfo(userId int) (models.UserInfo, error) {
	var user models.UserInfo

	query := fmt.Sprintf("SELECT name, email, phone, driver_license FROM %s WHERE id=$1", "users")
	err := r.db.Get(&user, query, userId)
	
	return user, err
}

func (r *UserPostgres) UpdateData(userId int, input models.UserUpdate) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValue = append(setValue, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Email != nil {
		setValue = append(setValue, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}
	if input.Phone != nil {
		setValue = append(setValue, fmt.Sprintf("phone=$%d", argId))
		args = append(args, *input.Phone)
		argId++
	}
	if input.DriverLicense != nil {
		setValue = append(setValue, fmt.Sprintf("driver_license=$%d", argId))
		args = append(args, *input.DriverLicense)
		argId++
	}
	if len(setValue) == 0 {
		return errors.New("No change")
	}
	setQuery := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s", "users", setQuery)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	
	return nil
}
