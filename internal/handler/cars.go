package handler

import (
	"CarRentalService/internal/models"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// addCar godoc
// @Summary Добавить автомобиль
// @Description Добавление нового автомобиля в систему (требуется авторизация)
// @Tags cars
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param input body models.Car true "Данные автомобиля"
// @Success 200 {object} map[string]int "ID созданного автомобиля"
// @Failure 400 {object} errorResponce "Неверный формат данных"
// @Failure 401 {object} errorResponce "Не авторизован"
// @Failure 500 {object} errorResponce "Ошибка сервера"
// @Router /api/car [post]
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

// getAllCars godoc
// @Summary Получить все автомобили
// @Description Получение списка всех автомобилей в системе
// @Tags cars
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} getAllCarsResponce "Список автомобилей"
// @Failure 500 {object} handler.errorResponce "Ошибка сервера"
// @Router /api/car [get]
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

// getCarById godoc
// @Summary Получить автомобиль по ID
// @Description Получение информации об автомобиле по его идентификатору
// @Tags cars
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "ID автомобиля"
// @Success 200 {object} models.Car "Данные автомобиля"
// @Failure 400 {object} handler.errorResponce "Неверный формат ID"
// @Failure 404 {object} handler.errorResponce "Автомобиль не найден"
// @Failure 500 {object} handler.errorResponce "Ошибка сервера"
// @Router /api/car/{id} [get]
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

// deleteCar godoc
// @Summary Удалить автомобиль
// @Description Удаление автомобиля по ID (требуется авторизация владельца)
// @Tags cars
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth 
// @Param id path int true "ID автомобиля"
// @Success 200 {object} statusResponce "Статус операции"
// @Failure 400 {object} handler.errorResponce "Неверный формат ID"
// @Failure 401 {object} handler.errorResponce "Не авторизован"
// @Failure 404 {object} handler.errorResponce "Автомобиль не найден или нет прав"
// @Failure 500 {object} handler.errorResponce "Ошибка сервера"
// @Router /api/car/{id} [delete]
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