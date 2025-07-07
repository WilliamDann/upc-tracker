package shared

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// get the authed user from request headers
func GetAuthorizedUser(r *http.Request) (*string, bool) {
	header := r.Header.Get("Authorization")

	// try to read authorization header
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, false
	}

	// get user assoicated
	token, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, false
	}

	// validate token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Validate expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, false
			}
		}

		// return subject of claim
		val, err := claims.GetSubject()
		if err != nil {
			return nil, false
		}
		return &val, true
	}

	return nil, false
}
