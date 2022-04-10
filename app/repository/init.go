package repository

import (
	"app/main/model"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

)

var db *gorm.DB

// DB接続
func Init() {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_DATABASE")
	PROTOCOL := "tcp(db:3306)"
	CONNECT := fmt.Sprintf("%s:%s@%s/%s?parseTime=true",
		USER, PASS, PROTOCOL, DBNAME)
	err := errors.New("")
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	autoMigrate()
}

// マイグレーション
func autoMigrate() {
	db.AutoMigrate(&model.Book{})
}

// DB接続解除
func Close() {
	db.Close()
}
