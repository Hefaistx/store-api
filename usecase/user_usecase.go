package usecase

import (
	"tokocikbosapi/model"
	"tokocikbosapi/repository"
)

type UserUsecase interface {
	FindUserByUsernamePassword(username, password string) (model.UserCredential, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func (uc *userUsecase) FindUserByUsernamePassword(username, password string) (model.UserCredential, error) {
	return uc.repo.FindUserByUsernamePasswordQuery(username, password)
}
