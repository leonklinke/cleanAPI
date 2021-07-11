package usecases

import (
	"cleanApi/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Show(userId int) (*domain.User, error) {
	return interactor.UserRepository.FindById(userId)
}
