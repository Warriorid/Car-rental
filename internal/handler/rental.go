package handler

import (
	"CarRentalService/internal/models"
	"CarRentalService/internal/repository"
	"CarRentalService/internal/service"
	"database/sql"
	"errors"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func(h *Handler) startRental(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input models.StartRent
	if err := c.BindJSON(&input); err != nil {
		newErrorRessponce(c, http.StatusBadRequest, "invalid data")
		return
	}
	rentId, err := h.service.StartRental(userId, input)
	if err != nil {
		if err == service.ErrInvalidDateRange || err == service.ErrParseDate {
			newErrorRessponce(c, http.StatusBadRequest, err.Error())
			return
		} else if errors.Is(err, sql.ErrNoRows) {
			newErrorRessponce(c, http.StatusNotFound, "car not found")
			return
		}
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]int {
		"id": rentId,
	})
}

func (h *Handler) endRendal(c *gin.Context) {
	rentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorRessponce(c, http.StatusBadRequest, "invalid item id param")
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	price, err := h.service.EndRental(rentId, userId)
	if err != nil {
		if errors.Is(err, repository.ErrRentalNotFound) {
			newErrorRessponce(c, http.StatusNotFound, "rental not found")
			return
		} else if errors.Is(err, repository.ErrRentalCompleted){
			newErrorRessponce(c, http.StatusNotFound, "rental already complate")
			return
		}
		if errors.Is(err, service.ErrPermission) {
			newErrorRessponce(c, http.StatusForbidden, err.Error())
			return
		}
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string] float64 {
		"total price": price,
	})

}

func (h *Handler) rentalHistory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	

	rentalsList, err := h.service.Rental.RentalHistory(userId)
	if err != nil {
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllRentalsResponse{
		Data: rentalsList,
	})
}