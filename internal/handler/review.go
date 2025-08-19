package handler

import (
	"CarRentalService/internal/models"
	"CarRentalService/internal/repository"
	"CarRentalService/internal/service"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) leaveReview(c *gin.Context){
	var input models.Reviews
	if err := c.BindJSON(&input); err != nil {
		newErrorRessponce(c, http.StatusBadRequest, "invalid data")
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	reviewId, err := h.service.Reviews.LeaveReview(userId, input)
	if err != nil {
		if err == service.ErrReviewPermission {
			newErrorRessponce(c, http.StatusForbidden, err.Error())
			return
		} else if errors.Is(err, repository.ErrRentalNotFound){
			newErrorRessponce(c, http.StatusNotFound, err.Error())
			return
		} else if errors.Is(err, service.ErrReviewAlreadySent) {
			newErrorRessponce(c, http.StatusConflict, err.Error())
			return
		}
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"id": reviewId,
	})
}