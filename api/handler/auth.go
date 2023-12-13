package handler

import (
	"context"
	"fmt"
	"geo/config"
	"geo/models"
	"geo/pkg/helper"
	"geo/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) SignUp(c *gin.Context) {

	var user models.CreateUserReq
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	hashedPass, err := helper.GeneratePasswordHash(user.Password)
	if err != nil {
		h.log.Error("error while generating hash password:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	user.Password = string(hashedPass)

	resp, err := h.storage.User().CreateUser(c.Request.Context(), &user)
	if err != nil {
		fmt.Println("error User Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "username is already used, enter another one")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created", "id": resp})
}

func (h *Handler) Login(c *gin.Context) {
	var req models.LoginUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fields in body"})
		return
	}

	resp, err := h.storage.User().GetUserByEmail(context.Background(), &models.LoginUserReq{
		Email: req.Email,
	})
	if err != nil {
		h.log.Error("error get by email:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not found email"})
		return
	}

	// Compare hashed password with plain text password
	err = helper.ComparePasswords([]byte(resp.Password), []byte(req.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "login or password didn't match"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "password comparison failed"})
		}
		return
	}

	m := make(map[string]interface{})
	m["username"] = resp.Username
	m["email"] = resp.Email

	token, _ := helper.GenerateJWT(m, config.TokenExpireTime, config.JWTSecretKey)

	c.SetCookie("jwt", token, 60*60*24, "/", "localhost", false, true)

	c.JSON(http.StatusOK, models.LoginUserRes{
		ID:          int(resp.ID),
		Username:    resp.Username,
		AccessToken: token,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
