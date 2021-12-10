package handlers

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"user-plus/domain/errs"
	"user-plus/endpoints/dto/response"
	"user-plus/endpoints/handlers"
	"user-plus/tests/endpoints/utils"
	IMockInterfaces "user-plus/tests/mocks/services"
)

func TestFindUserByEmail(t *testing.T) {
	t.Run("HandlersFindUserByEmailPass", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		expectUser := &response.UserReponse{
			Id:    1,
			Name:  "carlinhos",
			Email: "micledson@gmail.com",
		}

		userService := IMockInterfaces.NewMockIUserService(controller)
		userService.EXPECT().FindUserByEmail(gomock.Any()).Return(expectUser, nil)
		handler := handlers.NewUserHandlerFromService(userService)

		expectJson, err := json.Marshal(expectUser)
		assert.Nil(t, err)

		utils.RunGenericMethodGet(t,
			200,
			handler.FindUserByEmail,
			"/send/:email/find",
			utils.NewCtxParams().Add("email", "micledson@gmail.com"),
			string(expectJson),
		)
	})
	t.Run("HandlersFindUserByEmailFail", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		expectError := &errs.ApiErr{
			StatusCode: 400,
			Message:    "Dados solicitados não encontrados.",
		}

		userService := IMockInterfaces.NewMockIUserService(controller)
		userService.EXPECT().FindUserByEmail(gomock.Any()).Return(nil, expectError)
		handler := handlers.NewUserHandlerFromService(userService)

		utils.RunGenericMethodGet(t,
			http.StatusBadRequest,
			handler.FindUserByEmail,
			"/send/:email/find",
			utils.NewCtxParams().Add("email", "micledson@gmail.com"),
			utils.GetJsonNoData(),
		)
	})
}
