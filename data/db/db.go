package db

import (
	"fmt"
	"log"
	"time"
	"user-plus/domain/errs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetConnection() (*sqlx.DB, *errs.ApiErr) {
	// Load vars env
	//user := utils.GetVar("DB_USER")
	//pwd := utils.GetVar("DB_PASSWORD")
	//name := utils.GetVar("DB_NAME")
	//host := utils.GetVar("DB_ADDRESS")
	//port := utils.GetVar("DB_PORT")

	user := "root"
	pwd := "12345678"
	name := "userplus"

	url := fmt.Sprintf("%s:%s@/%s", user, pwd, name)
	// init database
	db, err := sqlx.Open("mysql", url)

	if err != nil {
		err := fmt.Sprintf(`Error while accessing database: %s`, err.Error())
		apiErr := &errs.ApiErr{Message: err}
		log.Print(apiErr)
		return nil, apiErr
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
