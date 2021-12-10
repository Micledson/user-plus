package interfaces

import (
	"user-plus/endpoints/dto/response"
)

type IUserService interface {
	FindUserByEmail(email string) (*response.UserReponse, error)
}
