package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitDB(dataSourceName string)(*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	//Ping = check database availability
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db,nil
}
