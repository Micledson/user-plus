package services

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"user-plus/domain"
	"user-plus/domain/errs"
	"user-plus/endpoints/dto/response"
	"user-plus/services"
	IMockInterfaces "user-plus/tests/mocks/repository"
)

func TestFindUserByEmail(t *testing.T) {
	t.Run("ServicesFindUserByEmailPass", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		userRepo := IMockInterfaces.NewMockIUserRepository(controller)

		userDomain := &domain.User{
			Id:    1,
			Name:  "Micledson",
			Email: "micledson@gmail.com",
		}

		userRepo.EXPECT().FindUserByEmail(gomock.Any()).Return(userDomain, nil)

		service := services.NewUserService(userRepo)

		result, err := service.FindUserByEmail("micledson@gmail.com")

		expect := &response.UserReponse{
			Id:    1,
			Name:  "Micledson",
			Email: "micledson@gmail.com",
		}

		assert.Equal(t, expect, result)
		assert.Nil(t, err)
	})
	t.Run("ServicesFindUserByEmailFail", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		userRepo := IMockInterfaces.NewMockIUserRepository(controller)

		errExpected := &errs.ApiErr{
			StatusCode: 400,
			Message:    "Dados Solicitados não encontrados",
		}

		userRepo.EXPECT().FindUserByEmail(gomock.Any()).Return(nil, errExpected)

		service := services.NewUserService(userRepo)

		_, err := service.FindUserByEmail("micledson@gmail.com")

		assert.Error(t, err)

	})
}
