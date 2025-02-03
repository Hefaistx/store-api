package modelutil

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtPayloadClaims struct {
	jwt.RegisteredClaims
	CredId int
	Role   string
}
