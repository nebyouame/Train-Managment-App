package service

import (

	"github.com/travel/courses/book"
	"github.com/travel/courses/model"
)

type BookService struct {
	bookRepo book.BookRepostory
}

func NewBookService(bookRepository book.BookRepostory) book.BookService {
	return &BookService{bookRepo:bookRepository}
}

func (bs *BookService) Books() ([]model.Book, []error) {
	bks, errs := bs.bookRepo.Books()
	if len(errs) > 0 {
		return nil, errs
	}
	return bks, errs
}

func (bs *BookService) Book(id uint) (*model.Book, []error){
	bk, errs := bs.bookRepo.Book(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}

func (bs *BookService) CustomerBooks(customer *model.User) ([]model.Book, []error){
	bks, errs :=bs.bookRepo.CustomerBooks(customer)
	if len(errs) > 0 {
		return nil, errs
	}
	return bks,errs
}

func (bs *BookService) UpdateBook(book *model.Book) (*model.Book, []error){
	bk, errs := bs.bookRepo.UpdateBook(book)
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}

func (bs *BookService) DeleteBook(id uint) (*model.Book, []error){
	bk, errs := bs.DeleteBook(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}

func (bs *BookService) StoreBook(book *model.Book) (*model.Book, []error){
	bk, errs := bs.StoreBook(book)
	if len(errs) > 0 {
		return nil, errs
	}
	return bk, errs
}

