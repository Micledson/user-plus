package interfaces

import (
	"user-plus/domain"
	"user-plus/domain/errs"
)

type IUserRepository interface {
	FindUserByEmail(email string) (*domain.User, *errs.ApiErr)
}
