package repository

import (

	"github.com/travel/courses/model"
	"github.com/travel/courses/schedule"
	"github.com/jinzhu/gorm"
)

type BookGormRepo struct {
	conn *gorm.DB
}

func NewBookGormRepo(db *gorm.DB) schedule.BookRepository {
	return &BookGormRepo{conn:db}
}

func (bRepo *BookGormRepo) Books() ([]model.Book, []error) {
	zss := []model.Book{}
	errs := bRepo.conn.Find(&zss).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return zss, nil
}

func (bRepo *BookGormRepo) Book(id uint) (*model.Book, []error) {
	zs := model.Book{}
	errs := bRepo.conn.First(&zs, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &zs, nil
}

func (bRepo *BookGormRepo) UpdateBook(book *model.Book) (*model.Book, []error){
	bat := book
	errs := bRepo.conn.Save(bat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return bat, nil
}

func (bRepo *BookGormRepo) DeleteBook(id uint) (*model.Book, []error) {
	bat, errs := bRepo.Book(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = bRepo.conn.Delete(bat, bat.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return bat, nil
}

func (bRepo *BookGormRepo) StoreBook(book *model.Book) (*model.Book, []error) {
	bat := book
	errs := bRepo.conn.Create(bat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return bat, nil
}


