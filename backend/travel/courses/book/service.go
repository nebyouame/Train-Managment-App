package book

import (

	"github.com/travel/courses/model"
)

type BookService interface {
	Books() ([]model.Book, []error)
	Book(id uint) (*model.Book, []error)
	CustomerBooks(customer *model.User) ([]model.Book, []error)
	UpdateBook(book *model.Book) (*model.Book, []error)
	DeleteBook(id uint) (*model.Book, []error)
	StoreBook(book *model.Book) (*model.Book, []error)
}
