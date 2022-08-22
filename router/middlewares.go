package router

import (
	"github.com/devazizi/go-crm/contract/response"
	infra "github.com/devazizi/go-crm/infrastructure"
	"github.com/devazizi/go-crm/repository"
	"github.com/devazizi/go-crm/service/auth"
	"github.com/devazizi/go-crm/service/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func middlewareCORS() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func middlewareCheckAuthenticated(DB infra.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		authorizationKey := c.GetHeader("Authorization")

		if authorizationKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{Status: false, Message: "authorization fail"})
			return
		}

		tokenData := strings.Split(authorizationKey, " ")

		jwtCredentials, _ := jwt.GetClaimsFromToken(tokenData[1])
		userInfo := jwtCredentials["UserInfo"]
		user := userInfo.(map[string]interface{})

		userId := user["id"].(float64)
		tokenHash, isCorrect := user["token"].(string)

		if !isCorrect {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{Status: false, Message: "authentication fail"})
			return
		}

		isValid := repository.ValidateTokenExistInStorage(DB, tokenHash, int(userId))

		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{Status: false, Message: "authentication fail"})
			return
		}

		authentication := auth.Authentication{UserId: int(userId), Token: tokenData[1]}
		authentication.SetAuthentication()

		c.Next()
	}
}
