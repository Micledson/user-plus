package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"user-plus/data/repository"
	"user-plus/domain"
	"user-plus/tests"
)

func TestFindUserByEmail(t *testing.T) {
	t.Run("RepositoryFindUserByEmailPass", func(t *testing.T) {
		tests.SetUp(insertUser()...)

		repository := repository.NewUserRepository()

		email := "micledson@gmail.com"
		result, err := repository.FindUserByEmail(email)

		expected := &domain.User{
			Id:    1,
			Name:  "Micledson",
			Email: email,
		}

		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})
	t.Run("RepositoryFindUserByEmailFail", func(t *testing.T) {
		tests.SetUp("TRUNCATE TABLE user;")

		repository := repository.NewUserRepository()

		email := "micledson@gmail.com"
		_, err := repository.FindUserByEmail(email)

		assert.Error(t, err)
	})
}

func insertUser() []string {
	return []string{
		"TRUNCATE TABLE user;",
		"INSERT INTO user VALUES (1, 'Micledson', 'micledson@gmail.com', 'password');",
		"INSERT INTO user VALUES (2, 'Carlos', 'carlosemail@gmail.com', 'password');",
	}
}
