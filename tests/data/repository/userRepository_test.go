package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"user-plus/data/repository"
	"user-plus/domain"
	"user-plus/tests"
)

func TestFindUserByEmail(t *testing.T) {

	tests.SetUp("TRUNCATE TABLE user;",
		"INSERT INTO user VALUES (1, 'Micledson', 'micledson14@gmail.com', 'password');",
	)

	repository := repository.NewUserRepository()

	email := "micledson14@gmail.com"
	result, err := repository.FindUserByEmail(email)

	expected := &domain.User{
		Id:    1,
		Name:  "Micledson",
		Email: email,
	}

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}
