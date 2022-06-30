package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id int, email string, role int, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add((30 * time.Minute)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["email"] = email
	claims["id"] = id
	claims["role"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
