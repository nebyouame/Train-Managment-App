package service

import (

	"github.com/travel/courses/model"
	"github.com/travel/courses/schedule"
)

type Bookservice struct {
	bookRepo schedule.BookRepository
}

func NewBookService(bookRepository schedule.BookRepository) schedule.BookService{
	return &Bookservice{bookRepo:bookRepository}
}

func (bs *Bookservice) Books() ([]model.Book, []error) {
	books, errs := bs.bookRepo.Books()
	if len(errs) > 0 {
		return nil, errs
	}

	return books, nil
}

func (bs *Bookservice) StoreBook(book *model.Book) (*model.Book, []error){
	bk, errs := bs.bookRepo.StoreBook(book)
	if len(errs) > 0 {
		return nil, errs
	}

	return bk, nil
}

func (bs *Bookservice) Book(id uint) (*model.Book, []error) {
	b, errs := bs.bookRepo.Book(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return b, nil
}

func (bs *Bookservice) UpdateBook(book *model.Book) (*model.Book, []error){
	bak, errs := bs.bookRepo.UpdateBook(book)
	if len(errs) > 0 {
		return nil, errs
	}

	return bak, nil
}

func (bs *Bookservice) DeleteBook(id uint) (*model.Book, []error)  {
	bak, errs := bs.bookRepo.DeleteBook(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return bak, nil
}

