package handler

import (
	"CarRentalService/internal/models"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getInfo godoc
// @Summary Получить информацию о пользователе
// @Description Получение данных текущего авторизованного пользователя
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.User "Данные пользователя"
// @Failure 401 {object} errorResponce "Не авторизован"
// @Failure 500 {object} errorResponce "Ошибка сервера"
// @Router /api/user [get]
func (h *Handler) getInfo(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	user, err := h.service.User.GetInfo(userId)
	if err != nil {
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}


// updateInfo godoc
// @Summary Обновить информацию пользователя
// @Description Обновление данных текущего авторизованного пользователя
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param input body models.UserUpdate true "Данные для обновления"
// @Success 200 {object} statusResponce "Статус операции"
// @Failure 400 {object} errorResponce "Неверный формат данных"
// @Failure 401 {object} errorResponce "Не авторизован"
// @Failure 500 {object} errorResponce "Ошибка сервера"
// @Router /api/user [put]	
func (h *Handler) updateInfo(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input models.UserUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorRessponce(c, http.StatusBadRequest, "uncorrected data")
		return
	}
	if err := h.service.User.UpdateData(userId, input); err != nil {
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}

// deleteUser godoc
// @Summary Удалить пользователя
// @Description Удаление аккаунта текущего авторизованного пользователя
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} statusResponce "Статус операции"
// @Failure 401 {object} errorResponce "Не авторизован"
// @Failure 404 {object} errorResponce "Пользователь не найден"
// @Failure 500 {object} errorResponce "Ошибка сервера"
// @Router /api/user [delete]
func (h *Handler) deleteUser(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	if err := h.service.User.Delete(userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newErrorRessponce(c, http.StatusNotFound, "user not found")
			return
		}
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.addTokenToBlackList(c)
	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}
