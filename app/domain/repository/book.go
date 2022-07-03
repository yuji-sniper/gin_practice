package repository

import "app/main/domain/model"

type BookContract interface {
	FetchBooks() ([]model.Book)
	FindBook(id int) (model.Book)
	CreateBook(book *model.Book)
	UpdateBook(id int, book *model.Book)
	DeleteBook(id int)
}
