package models

type Car struct {
	Id int `json:"-" db:"id"`
	Model string `json:"model" binding:"required"`
	Year int `json:"year" binding:"required"`
	Color string `json:"color" binding:"required"`
	Mileage int `json:"mileage" binding:"required"`
	PricePerDay float64 `json:"price_per_day" binding:"required"`
	IsAvailable bool `json:"is_available"`
	Location string `json:"location" binding:"required"`
	OwnerId int `json:"-" db:"owner_id"`
}

type BriefCarInfo struct {
	Model string `json:"model" db:"model"`
	Color string `json:"color" db:"color"`
	PricePerDay float64 `json:"price_per_day" db:"price_per_day"`
}

type GetCarInfo struct {
	Model string `json:"model" binding:"required"`
	Year int `json:"year" binding:"required"`
	Color string `json:"color" binding:"required"`
	Mileage int `json:"mileage" binding:"required"`
	PricePerDay float64 `json:"price_per_day" binding:"required"`
	Location string `json:"location" binding:"required"`
	OwnerName string `json:"owner_name"`
}