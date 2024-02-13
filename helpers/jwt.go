package helpers

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var secretKey string

func init() {
	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
  }
	secretKey = os.Getenv("SECRET")
}

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}	
	stringToken := strings.Split(headerToken, " ")[1]
	
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {		
		return nil, errResponse
	}	
	
	return token.Claims.(jwt.MapClaims), nil
}