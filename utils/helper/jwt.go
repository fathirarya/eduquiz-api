package helper

import (
	"eduquiz-api/model/web"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userLoginResponse *web.UserLoginResponse, id uint) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["fullname"] = userLoginResponse.Fullname
	claims["email"] = userLoginResponse.Email
	claims["roles"] = userLoginResponse.Roles
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}
