package repository

import (
	"github.com/labstack/gommon/log"
	"user-plus/data/db"
	"user-plus/domain"
	"user-plus/domain/errs"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur UserRepository) FindUserByEmail(email string) (*domain.User, *errs.ApiErr) {
	conn, err := db.GetConnection()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer conn.Close()

	var user domain.User
	error := conn.Get(&user, `SELECT id, name, email FROM user where email = ?;`, email)
	if error != nil {
		log.Error(error)
		return nil, &errs.ApiErr{Message: "Dados solicitados não encontrados."}
	}

	return &user, nil

}
