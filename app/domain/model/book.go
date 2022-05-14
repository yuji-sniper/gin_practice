package model

import "time"

type Book struct {
	Id int `form:"id" gorm:"AUTO_INCREMENT"`
	Title string `form:"title"`
	Content string `form:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
