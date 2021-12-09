package db

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetConnection() (*sqlx.DB, error) {
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
		log.Print("Error while accessing database: " + err.Error())
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
