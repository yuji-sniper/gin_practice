package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db := DBConnect()
	defer db.Close()
}

func DBConnect() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_DATABASE")
	PROTOCOL := "tcp(db:3306)"
	CONNECT := fmt.Sprintf("%s:%s@%s/%s?parseTime=true",
		USER, PASS, PROTOCOL, DBNAME)
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return db
}
