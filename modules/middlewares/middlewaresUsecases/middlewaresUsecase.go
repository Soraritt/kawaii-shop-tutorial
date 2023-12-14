package middlewaresUsecases

import (
	"github.com/soraritt/kawaii-shop-tutorial/modules/middlewares"
	"github.com/soraritt/kawaii-shop-tutorial/modules/middlewares/middlewaresRepositories"
)

type IMidderwaresUsecase interface {
	FindAccessToken(userId, accessToken string) bool
	FindRole() ([]*middlewares.Role, error)
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMidderwaresRepository
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepositories.IMidderwaresRepository) IMidderwaresUsecase {

	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}

}

func (u *middlewaresUsecase) FindAccessToken(userId, accessToken string) bool {
	return u.middlewaresRepository.FindAccessToken(userId, accessToken)
}

func (u *middlewaresUsecase) FindRole() ([]*middlewares.Role, error) {
	roles, err := u.middlewaresRepository.FindRole()
	if err != nil {
		return nil, err
	}
	return roles, nil
}
