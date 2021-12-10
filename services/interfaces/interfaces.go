package interfaces

import (
	"user-plus/domain/errs"
	"user-plus/endpoints/dto/response"
)

type IUserService interface {
	FindUserByEmail(email string) (*response.UserReponse, *errs.ApiErr)
}
