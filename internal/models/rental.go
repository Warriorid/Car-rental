package models

import "time"

type Rental struct{
	Id int `json:"-" db:"id"`
	CarId int `json:"car_id" db:"car_id" binding:"required"`
	UserId int `json:"user_id" db:"user_id"`
	StartDate time.Time `json:"start_date" db:"start_date" binding:"required"`
	EndDate time.Time `json:"end_date" db:"end_date" binding:"required"`
	TotalPrice float64 `json:"total_price" db:"total_price"`
	Status string `json:"status" db:"status"`
}

type RentalHistory struct {
	Car string `json:"car"`
	User string `json:"user"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate time.Time `json:"end_date" db:"end_date"`
	TotalPrice float64 `json:"total_price" db:"total_price"`
	Status string `json:"status" db:"status"`
}


type StartRent struct {
    CarId     int    `json:"car_id" binding:"required"`
    StartDate string `json:"start_date" binding:"required"`
    EndDate   string `json:"end_date" binding:"required"`
}