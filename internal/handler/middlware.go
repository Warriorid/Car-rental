package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)


const (
	userCtx = "userId"
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity (c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorRessponce(c, http.StatusUnauthorized, "header is empty")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorRessponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	token := headerParts[1]
	defer cancel()
	exists, err := h.redisClient.Exists(ctx, "blacklist:"+ token).Result()
	if err != nil {
		newErrorRessponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	if exists == 1 {
		newErrorRessponce(c, http.StatusUnauthorized, "token in black list")
		return
	}
	userId, err := h.service.Autorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorRessponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error){
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorRessponce(c, http.StatusUnauthorized, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorRessponce(c, http.StatusUnauthorized, "invalid type id")
		return 0, errors.New("invalid type id")
	}
	return idInt, nil
}

func (h *Handler) addTokenToBlackList(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
    if header == "" {
        newErrorRessponce(c, http.StatusBadRequest, "empty auth header")
        return
    }
    
    headerParts := strings.Split(header, " ")
    if len(headerParts) != 2 {
        newErrorRessponce(c, http.StatusBadRequest, "invalid auth header format")
        return
    }
	token := headerParts[1]
	expiresAt := time.Now().Add(time.Hour*24)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	h.redisClient.Set(ctx, "blacklist:" + token, "1", time.Until(expiresAt))
	
}