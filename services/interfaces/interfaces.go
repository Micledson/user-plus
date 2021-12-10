package interfaces

import "user-plus/tests/endpoints/dto/response"

type IUserService interface {
	FindUserByEmail(email string) (*response.UserReponse, error)
}
