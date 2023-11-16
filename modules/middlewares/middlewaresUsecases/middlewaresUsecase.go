package middlewaresUsecases

import "github.com/soraritt/kawaii-shop-tutorial/modules/middlewares/middlewaresRepositories"

type IMidderwaresUsecase interface {
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMidderwaresRepository
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepositories.IMidderwaresRepository) IMidderwaresUsecase {

	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}
