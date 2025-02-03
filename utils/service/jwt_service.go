package service

import (
	"fmt"
	"time"
	"tokocikbosapi/config"
	"tokocikbosapi/model"
	modelutil "tokocikbosapi/utils/model_util"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(user model.UserCredential) (string, error)
	VerifyToken(tokenString string) (modelutil.JwtPayloadClaims, error)
}

type jwtService struct {
	cfg config.TokenConfig
}

func (j *jwtService) GenerateToken(user model.UserCredential) (string, error) {
	//get the secret key
	tokenKey := j.cfg.JwtSignatureKey

	//fill the claims
	claims := modelutil.JwtPayloadClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.ApplicationName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.AccesTokenLifeTime)),
		},
		UserId: user.ID,
		Role:   user.Roles,
	}

	//token
	jwtNewClaims := jwt.NewWithClaims(j.cfg.JwtSigningMethod, claims)
	token, err := jwtNewClaims.SignedString(tokenKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *jwtService) VerifyToken(tokenString string) (modelutil.JwtPayloadClaims, error) {
	//parsing token
	tokenParse, err := jwt.ParseWithClaims(tokenString, &modelutil.JwtPayloadClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return modelutil.JwtPayloadClaims{}, err
	}

	claim, ok := tokenParse.Claims.(*modelutil.JwtPayloadClaims)
	if !ok {
		return modelutil.JwtPayloadClaims{}, fmt.Errorf("claim error")
	}
	return *claim, nil
}

func NewJwtService(cfg config.TokenConfig) JwtService {
	return &jwtService{cfg: cfg}
}
