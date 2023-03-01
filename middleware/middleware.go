package middleware

import (
	"fmt"
	"github.com/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/response"
	"net/http"
	"strings"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if error := validateToken(r); error != nil {
			response.Error(w, http.StatusUnauthorized, error)
		} else {
			next(w, r)
		}
	}
}

func validateToken(r *http.Request) error {
	_, error := jwt.Parse(extractToken(r), getSecretKey)

	if error != nil {
		return error
	}

	return nil
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	splitedBearerToken := strings.Split(bearerToken, " ")

	if len(splitedBearerToken) == 2 {
		return splitedBearerToken[1]
	} else {
		return ""
	}

}

func getSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("wrong assignature method. %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
