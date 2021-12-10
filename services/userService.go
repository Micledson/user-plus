package services

import (
	"github.com/labstack/gommon/log"
	"user-plus/data/interfaces"
	"user-plus/domain/errs"
	"user-plus/endpoints/dto/response"
)

type UserService struct {
	repostisory interfaces.IUserRepository
}

func NewUserService(repo interfaces.IUserRepository) *UserService {
	return &UserService{repostisory: repo}
}

func (us UserService) FindUserByEmail(email string) (*response.UserReponse, *errs.ApiErr) {
	user, err := us.repostisory.FindUserByEmail(email)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	userResponse := user.ConvertToDTO()

	return userResponse, nil
}
