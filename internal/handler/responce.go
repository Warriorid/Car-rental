package handler

import (
	"CarRentalService/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponce struct {
	Message string `json:"message"`
}

func newErrorRessponce(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponce{message})
}

type statusResponce struct {
	Status string `json:"status"`
}

type getAllCarsResponce struct{
	Data []models.BriefCarInfo
}

type getAllRentalsResponse struct{
	Data []models.RentalHistory
}