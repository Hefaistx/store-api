package usecase

import (
	"tokocikbosapi/utils/service"
)

type AuthenticationUsecase interface {
	Login(username string, password string) (string, error)
}

type authenticationUsecase struct {
	userUC     UserUsecase
	jwtService service.JwtService
}

func (uc *authenticationUsecase) Login(username string, password string) (string, error) {
	cred, err := uc.userUC.FindUserByUsernamePassword(username, password)
	if err != nil {
		return "", err
	}
	token, err := uc.jwtService.GenerateToken(cred)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewAuthenticationUsecase(userUC UserUsecase, jwtService service.JwtService) AuthenticationUsecase {
	return &authenticationUsecase{userUC: userUC, jwtService: jwtService}
}
