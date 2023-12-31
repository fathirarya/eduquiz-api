package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenStudent(StudentID uint) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"sub":  StudentID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"iat":  time.Now().Unix(),
		"role": "Student",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenTeacher(TeacherID uint) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"sub":  TeacherID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"iat":  time.Now().Unix(),
		"role": "Teacher",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
