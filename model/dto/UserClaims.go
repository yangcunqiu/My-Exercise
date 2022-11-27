package dto

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	Id   uint
	Name string
	jwt.StandardClaims
}
