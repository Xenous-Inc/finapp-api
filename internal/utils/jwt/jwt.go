package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(sessionId string, jwtSecret string) (string, error) {
	key := []byte(jwtSecret)
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["sessionId"] = sessionId 
    tokenString, err := token.SignedString(key)
    if err != nil {
        fmt.Println("Ошибка подписи JWT токена:", err)
        return "", err
    }
    fmt.Println("JWT токен успешно подписан")
    return tokenString, nil
}
