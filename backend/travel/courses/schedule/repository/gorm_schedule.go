package repository

import (
	"github.com/travel/courses/model"
	"github.com/travel/courses/schedule"
	"github.com/jinzhu/gorm"
)

type ScheduleGormRepo struct {
	conn *gorm.DB
}

func NewScheduleGormRepo(db *gorm.DB) schedule.ScheduleRepository {
	return &ScheduleGormRepo{conn:db}
}

func (sRepo *ScheduleGormRepo) Schedules() ([]model.Schedule, []error) {
	schs := []model.Schedule{}
	errs := sRepo.conn.Find(&schs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return schs, nil
}

func (sRepo *ScheduleGormRepo) Schedule(id uint) (*model.Schedule, []error) {
	sch := model.Schedule{}
	errs := sRepo.conn.First(&sch, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return &sch, nil
}

func (sRepo *ScheduleGormRepo) UpdateSchedule(schedule *model.Schedule) (*model.Schedule, []error){
	sat := schedule
	errs := sRepo.conn.Save(sat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (sRepo *ScheduleGormRepo) DeleteSchedule(id uint) (*model.Schedule, []error) {
	sat, errs := sRepo.Schedule(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = sRepo.conn.Delete(sat, sat.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (sRepo *ScheduleGormRepo) StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error) {
	sat := schedule
	errs := sRepo.conn.Create(sat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (sRepo *ScheduleGormRepo) InfosInSchedule(schedule *model.Schedule) ([]model.Info, []error) {
	infos := []model.Info{}
	sat, errs := sRepo.Schedule(schedule.ID)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = sRepo.conn.Model(sat).Related(&infos, "Infos").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return infos, nil
}
























