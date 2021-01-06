package auth

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
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

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(SIGNED_KEY), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
