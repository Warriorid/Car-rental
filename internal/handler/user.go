package handler

import (
	"CarRentalService/internal/models"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


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
