package tests

import (
	"github.com/labstack/echo/v4"
	"user-plus/data/db"
)

func SetUp(queries ...string) *echo.Echo {
	e := echo.New()
	conn, _ := db.GetConnection()
	for _, query := range queries {
		_, err := conn.Exec(query)
		if err != nil {
			panic(err)
		}
	}
	err := conn.Close()
	if err != nil {
		panic(err)
	}
	return e
}
