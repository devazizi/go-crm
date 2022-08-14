package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyJWTClaims struct {
	*jwt.RegisteredClaims
	UserInfo interface{}
}

var secret = []byte("secret key for my application it is just a secret")

func CreateToken(sub string, userInfo interface{}) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	exp := time.Now().Add(time.Hour * 24 * 30)
	token.Claims = &MyJWTClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   sub,
		},
		userInfo,
	}

	val, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return val, nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
