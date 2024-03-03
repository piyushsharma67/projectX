package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil

}

func VerifyToken(reqtoken string) (bool, string) {
	bearerToken := strings.Split(reqtoken, " ")
	if len(bearerToken) == 2 {
		claims := &CustomClaims{}
		token, _ := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid Token")
			}
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})
		if token.Valid {
			return true, claims.UserId
		}
		return false, ""
	}
	return false, ""
}