package controller

import (
	"github.com/devazizi/go-crm/contract/request"
	"github.com/devazizi/go-crm/contract/response"
	"github.com/devazizi/go-crm/infrastructure"
	"github.com/devazizi/go-crm/repository"
	"github.com/devazizi/go-crm/service/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginAPI(DB infrastructure.DB, req request.LoginRequest) gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{Message: err.Error(), Status: false})
			return
		}

		user, err := repository.CheckUserExistsByEmail(DB, req.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "email or password incorrect"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "email or password incorrect"})
			return
		}

		token, _ := jwt.CreateToken("2332-abcd-acdf-ccd2", user)

		c.JSON(http.StatusOK, response.Response{
			Status: response.SUCCESS, Data: response.LoginResponse{User: user, Token: token},
		})
	}
}

func RegisterAPI(DB infrastructure.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ForgetPassword(DB infrastructure.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
