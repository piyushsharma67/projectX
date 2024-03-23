package utils

import (
	"bytes"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(_password string) (string, error) {
	password := []byte(_password)
	if bytes.NewReader(password).Len() == 0 {
		return "", fmt.Errorf("password empty")
	}
	hash, _ := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	passhash := string(hash)
	return passhash, nil
}

func ComparePassword(hashedPassword, password string) error {
	var err error

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err == nil {
		return nil
	}

	return err
}

type CustomClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func CreateLoginJWTToken(email string, useragent string) (string, error) {

	claims:= CustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10000 * time.Hour).Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte("test"))
	if err != nil {
		return "", err
	}

	return token, nil

}

func VerifyToken(reqtoken string) (bool, string,string) {
	
	if len(reqtoken)!=0 {
		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(reqtoken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid Token")
			}
			return []byte("test"), nil
		})
		if token.Valid {
			return true, claims.UserId,claims.UserId
		}
		return false, err.Error(),""
	}
	return false, "token not provided",""
}