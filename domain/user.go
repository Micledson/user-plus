package domain

import (
	"user-plus/endpoints/dto/response"
)

type User struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (u User) ConvertToDTO() *response.UserReponse {
	return &response.UserReponse{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
	}
}
