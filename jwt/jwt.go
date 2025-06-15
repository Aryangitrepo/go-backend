package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

func CreateToken(username string) (string, error) {
	var secretKey string
	if err := loadEnv(); err.Error() == "Error loading .env fileopen .env: no such file or directory" {
		secretKey = "hello"
	}
	secretKey = os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "hello"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	loadEnv()
	secretKey := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
