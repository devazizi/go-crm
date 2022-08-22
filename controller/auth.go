package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/devazizi/go-crm/contract"
	"github.com/devazizi/go-crm/contract/mocks"
	"github.com/devazizi/go-crm/contract/request"
	"github.com/devazizi/go-crm/contract/response"
	"github.com/devazizi/go-crm/entity"
	"github.com/devazizi/go-crm/infrastructure"
	"github.com/devazizi/go-crm/repository"
	"github.com/devazizi/go-crm/service/auth"
	"github.com/devazizi/go-crm/service/helpers"
	"github.com/devazizi/go-crm/service/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginAPI(DB infrastructure.DB, requestValidator contract.ValidateLoginRequestFields) gin.HandlerFunc {

	return func(c *gin.Context) {
		loginRequest := request.LoginRequest{}
		if err := c.BindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{Message: err.Error(), Status: false})
			return
		}

		if err := requestValidator(loginRequest); err != nil {
			c.JSON(http.StatusUnprocessableEntity, response.Response{Message: err.Error(), Status: false})
			return
		}

		user, err := repository.CheckUserExistsByEmail(DB, loginRequest.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.Response{Message: "we can't match with or credential"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(loginRequest.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, response.Response{Message: "we can't match with or credential"})
			return
		}

		tokenHash := sha256.Sum256([]byte(helpers.RandomString(50)))
		token := hex.EncodeToString(tokenHash[:])
		_, err = repository.CreateClientToken(DB, user, token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "email or password incorrect"})
			return
		}

		userMock := mocks.UserMock{
			ID:    user.ID,
			Name:  user.Name,
			Token: token,
		}

		token, _ = jwt.CreateToken("2332-abcd-acdf-ccd2", userMock)

		c.JSON(http.StatusOK, response.Response{
			Status: response.SUCCESS, Data: response.LoginResponse{User: user, Token: token},
		})
	}
}

func RegisterAPI(DB infrastructure.DB, validator contract.ValidateRegisterRequestFields) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := request.RegisterRequest{}

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{Message: err.Error(), Status: false})
			return
		}

		if err := validator(request, DB); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{Message: err.Error(), Status: false})
			return
		}

		encryptedPass, _ := bcrypt.GenerateFromPassword([]byte(request.Paassword), 1)

		user := entity.User{
			Email:    request.Email,
			Name:     request.Name,
			Password: string(encryptedPass),
		}

		repository.CreateUser(DB, user)

		c.JSON(http.StatusCreated, response.Response{Status: response.SUCCESS, Data: request})
	}
}

func ForgetPassword(DB infrastructure.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ChangePassword(DB infrastructure.DB, validator contract.ValidateChangePasswordRequestFields) gin.HandlerFunc {

	return func(c *gin.Context) {
		request := request.ChangePasswordRequest{}

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{Message: err.Error(), Status: false})
			return
		}

		if err := validator(request); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{Message: err.Error(), Status: false})
			return
		}

		auth := auth.GetAuthentication()

		encryptedPass, _ := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 1)

		repository.New(DB).UpdateClientPassword(auth.UserId, string(encryptedPass))

		c.JSON(http.StatusOK, response.Response{Status: response.SUCCESS, Data: nil})
	}

}
