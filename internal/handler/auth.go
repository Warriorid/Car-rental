package handler

import (
	"CarRentalService/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"	
)



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