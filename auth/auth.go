package auth

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

const SIGNED_KEY = "abcdefghijklmn"

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1234567"
	claims["name"] = "test name"

	tokenString, _ := token.SignedString([]byte(SIGNED_KEY))

	w.Write([]byte(tokenString))
})
