package handler

import (
	"CarRentalService/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"	
)


// Register godoc
// @Summary Регистрация нового пользователя
// @Description Создает нового пользователя в системе
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.User true "Данные пользователя для регистрации"
// @Success 201 {object} errorResponce "Пользователь успешно создан"
// @Failure 400 {object} errorResponce "Неверный формат данных"
// @Failure 500 {object} errorResponce "Ошибка сервера"
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context){
	var input models.User
	err := c.BindJSON(&input)
	if err != nil {
		newErrorRessponce(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.service.Autorization.CreateUser(input)
	if err != nil {
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}


// Login godoc
// @Summary Авторизация пользователя
// @Description Вход в систему и получение токена
// @Tags auth
// @Accept  json
// @Produce  json
// @Param credentials body models.SignInInput true "Данные для входа"
// @Success 200 {object} errorResponce "Успешный вход, возвращает токен"
// @Failure 400 {object} errorResponce "Неверные учетные данные"
// @Failure 500 {object} errorResponce "Ошибка сервера"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorRessponce(c, http.StatusBadRequest, "uncorrected data")
		return
	}
	token, err := h.service.Autorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}