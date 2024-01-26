package jwtservice

import (
	"errors"
	"strings"

	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
	"github.com/golang-jwt/jwt/v5"
)

func NewToken(sessionId string, jwtSecret string) (string, error) {
	key := []byte(jwtSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sessionId"] = sessionId
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Error(err, "InternalServer, error signature JWT token", "jwtservice NewToken")
		return "", errors.New("Unathorized")
	}

	return tokenString, nil
}

func GetDecodeToken(tokenString string, jwtSecret string) (*jwt.Token, error) {
	if tokenString == "" {
		log.Warn("Authorization header is empty")
		return nil, errors.New("Unathorized")
	}

	tokenSlice := strings.Split(tokenString, "Bearer ")
	if len(tokenSlice) != 2 {
		log.Warn("Invalid Authorization header format")
		return nil, errors.New("Unathorized")
	}

	tokenString = tokenSlice[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, errors.New("Unathorized")
	}

	return token, err
}
