package models

import "time"

type Reviews struct {
	Id int `json:"-" db:"id"`
	RentalId int `json:"rental_id" db:"rental_id" binding:"required"`
	Rating int `json:"rating" db:"rating" binding:"required"`
	Comment string `json:"comment" db:"comment"`
	CreatedAt time.Time `json:"-"`
}