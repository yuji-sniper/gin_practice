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

// booksレコード１件取得
func (BookRepository) FindBook(id int) model.Book {
	book := model.Book{Id: id}
	db.First(&book)
	return book
}

// booksレコード更新
func (BookRepository) UpdateBook(id int, book *model.Book) {
	book.Id = id
	db.Save(&book)
}

// booksレコード１件削除
func (BookRepository) DeleteBook(id int) {
	book := model.Book{Id: id}
	db.Delete(&book)
}
