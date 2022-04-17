package model

import "time"

type Book struct {
	Id int `form:"id" gorm:"AUTO_INCREMENT"`
	Title string `form:"title"`
	Content string `form:"content"`
	CreatedAt time.Time `form:"craeted_at"`
	UpdatedAt time.Time `form:"updated_at"`
}