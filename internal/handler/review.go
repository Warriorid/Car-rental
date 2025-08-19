package handler

import (
	"CarRentalService/internal/models"
	"CarRentalService/internal/repository"
	"CarRentalService/internal/service"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

// leaveReview godoc
// @Summary Оставить отзыв
// @Description Создание отзыва для завершенной аренды (требуется авторизация)
// @Tags reviews
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param input body models.Reviews true "Данные отзыва"
// @Success 200 {object} map[string]int "ID созданного отзыва"
// @Failure 400 {object} errorResponce "Неверный формат данных"
// @Failure 401 {object} errorResponce "Не авторизован"
// @Failure 403 {object} errorResponce "Нет прав для отзыва"
// @Failure 404 {object} errorResponce "Аренда не найдена"
// @Failure 409 {object} errorResponce "Отзыв уже оставлен"
// @Failure 500 {object} errorResponce "Ошибка сервера"
// @Router /api/review [post]
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