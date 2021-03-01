package schedule

import (

	"github.com/travel/courses/model"
)

type ScheduleService interface {
	Schedules() ([]model.Schedule, []error)
	Schedule(id uint) (*model.Schedule, []error)
	UpdateSchedule(schedule *model.Schedule) (*model.Schedule, []error)
	DeleteSchedule(id uint) (*model.Schedule, []error)
	StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error)
	InfosInSchedule(schedule *model.Schedule) ([]model.Info, []error)
}



type BookService interface {
	Books() ([]model.Book, []error)
	Book(id uint) (*model.Book, []error)
	UpdateBook(book *model.Book) (*model.Book, []error)
	DeleteBook(id uint) (*model.Book, []error)
	StoreBook(book *model.Book) (*model.Book, []error)
}
