package utils

import (
	"My-Exercise/model/dto"
	"github.com/dgrijalva/jwt-go"
)

var signKey = []byte("111")

func GenerateToken(id uint, name string) (token string, err error) {
	userClaims := &dto.UserClaims{
		Id:             id,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	token, err = claims.SignedString(signKey)
	return
}

func ParseToken(tokenString string) (*dto.UserClaims, bool) {
	userClaims := new(dto.UserClaims)
	token, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return nil, false
	}
	if !token.Valid {
		return nil, false
	}
	if _, ok := token.Claims.(*dto.UserClaims); ok {
		return userClaims, true
	}
	return nil, false
}
