package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	stringConnection := "root:root@tcp(db:3306)/fabrica?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", stringConnection)

	if error != nil {
		return nil, error
	}
	
	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil
}
