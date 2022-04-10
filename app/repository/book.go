package repository

import "app/main/model"

type BookRepository struct{}

// books全件取得
func (BookRepository) FetchBooks() []model.Book {
	books := []model.Book{}
	db.Order("id desc").Find(&books)
	return books
}

// booksレコード１件作成
func (BookRepository) CreateBook(book *model.Book) {
	db.Create(&book)
}


