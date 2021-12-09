package interfaces

import "user-plus/domain"

type IUserRepository interface {
	FindUserByEmail(email string) (*domain.User, error)
}
