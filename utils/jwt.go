package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Create(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["roles"] = "ADMIN"
	claims["email"] = email
	claims["exp"] = time.Now().AddDate(0, 12, 0).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
func Extract(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	bearerTokenArr := strings.Split(bearerToken, " ")
	if len(bearerTokenArr) > 1 {
		return bearerTokenArr[1]
	}
	return ""

}

func ExtractEmailFromToken(c *gin.Context) (string, error) {
	var email string
	tokenString := Extract(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})

	if err != nil {
		return email, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email = fmt.Sprintf("%v", claims["email"])
	}

	return email, nil
}
func ExtractRolesFromToken(c *gin.Context) (string, error) {
	var roles string
	tokenString := Extract(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})

	if err != nil {
		return roles, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		roles = fmt.Sprintf("%v", claims["roles"])
	}

	return roles, nil
}
func Verify(c *gin.Context) error {
	tokenString := Extract(c)
	if tokenString != "" {
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_JWT")), nil
		})

		if err != nil {
			return err
		}
		return nil
	} else {
		return nil
	}
}
