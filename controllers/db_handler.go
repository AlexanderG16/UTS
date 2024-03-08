package controllers

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_uts?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func gormConnect() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: connect()}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
