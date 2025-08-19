package handler

import (
	"CarRentalService/internal/models"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (h *Handler) addCar(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var car models.Car
	if err := c.BindJSON(&car); err != nil {
		newErrorRessponce(c, http.StatusBadRequest, "uncorrect data")
		return
	}
	carId, err := h.service.Car.AddCar(userId, car)
	if err != nil {
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]int {
		"id": carId,
	})
}

func (h *Handler) getAllCars(c *gin.Context){

	carInfo, err := h.service.GetAllCars()
	if err != nil {
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllCarsResponce{
		Data: carInfo,
	})
}

func (h *Handler) getCarById(c *gin.Context) {
	carId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorRessponce(c, http.StatusBadRequest, "invalid car ID format")
		return
	}
	car, err := h.service.Car.GetCarById(carId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newErrorRessponce(c, http.StatusNotFound, "car not found")
			return
		} else {
			newErrorRessponce(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, car)
}

func (h *Handler) deleteCar(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	carId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorRessponce(c, http.StatusBadRequest, "invalid car ID format")
		return
	}
	err = h.service.Car.DeleteCar(userId, carId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newErrorRessponce(c, http.StatusNotFound, "car not found or you don't have permissions")
			return
		} else {
			newErrorRessponce(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
	
}