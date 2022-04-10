package model

import "time"

type Book struct {
	Id int `form:"id" gorm:"AUTO_INCREMENT"`
	Title string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
	CreatedAt time.Time `form:"craeted_at"`
	UpdatedAt time.Time `form:"updated_at"`
}