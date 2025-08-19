package models

type User struct{
	Id int `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	DriverLicense string `db:"driver_license" json:"driver_license" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfo struct {
	Name          string `json:"name" binding:"required"`
    Email         string `json:"email" binding:"required"`
    Phone         string `json:"phone" binding:"required"`
    DriverLicense string `db:"driver_license" json:"driver_license" binding:"required"`
}

type UserUpdate struct {
	Name *string `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	DriverLicense *string `db:"driver_license" json:"driver_license"`
}

type SignInInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}