package interfaces

import (
	"cleanApi/domain"
)

type UserRepository struct {
	DatabaseHandler DatabaseHandler
}

func (repository *UserRepository) FindById(userId int) (*domain.User, error) {
	return &domain.User{
		Id:   1,
		Name: "leon",
	}, nil
}
