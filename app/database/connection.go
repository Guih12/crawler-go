package database

import (
	"crawler/app/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

func Connect() (*sql.DB, error) {
	url, _ := utils.LoadingEnviroments()
	database, err := sql.Open("mysql", url)

	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		database.Close()
		return nil, err
	}

	return database, nil
}
