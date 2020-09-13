package model

import (
	"github.com/dgrijalva/jwt-go"
)

type JWT struct{
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}