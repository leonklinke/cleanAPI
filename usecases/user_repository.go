package usecases

import (
	"cleanApi/domain"
)

type UserRepository interface {
	FindById(int) (*domain.User, error)
}
