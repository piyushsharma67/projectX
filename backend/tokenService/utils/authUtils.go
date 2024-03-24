package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func CreateLoginJWTToken(email string,secret string) (string, error) {

	claims:= CustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10000 * time.Hour).Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil

}

func VerifyToken(reqtoken string,secret string) (bool, string,string) {
	
	if len(reqtoken)!=0 {
		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(reqtoken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid Token")
			}
			return []byte(secret), nil
		})
		if token.Valid {
			return true, claims.UserId,claims.UserId
		}
		return false, err.Error(),""
	}
	return false, "token not provided",""
}