package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

)

func GenerateNewAccessToken() (string, error) {
	secret := GetValue("JWT_SECRET_KEY")
	minuresCount, _ := strconv.Atoi(GetValue("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// create a JWT claim object
	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minuresCount)).Unix()

	// create a new JWT token with the JWT claim object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// convert the token in a string format
	t, err := token.SignedString([]byte(secret))

	// if conversion failed, return the error
	if err != nil {
		return "", err
	}

	// return the token
	return t, nil
}

