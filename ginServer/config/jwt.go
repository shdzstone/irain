package config

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

//const TokenExpireDuration = time.Hour * 2
const TokenExpireDuration = time.Second*10

var SecretKey = []byte("马云飘了")

