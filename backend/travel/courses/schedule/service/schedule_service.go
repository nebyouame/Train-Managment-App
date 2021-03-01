package service


import (
	"github.com/travel/courses/model"
	"github.com/travel/courses/schedule"
)

type ScheduleService struct {
	scheduleRepo schedule.ScheduleRepository
}

func NewScheduleService(SatRepo schedule.ScheduleRepository) schedule.ScheduleService {
	return &ScheduleService{scheduleRepo:SatRepo}
}

func (ss *ScheduleService) Schedules() ([]model.Schedule, []error) {
	schedules, errs := ss.scheduleRepo.Schedules()

	if len(errs) > 0 {
		return nil, errs
	}

	return schedules, nil
}

func (ss *ScheduleService) StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error){
	sat, errs := ss.scheduleRepo.StoreSchedule(schedule)
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (ss *ScheduleService) Schedule(id uint) (*model.Schedule, []error) {
	s, errs := ss.scheduleRepo.Schedule(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return s, nil
}

func (ss *ScheduleService) UpdateSchedule(schedule *model.Schedule) (*model.Schedule, []error){
	sat, errs := ss.scheduleRepo.UpdateSchedule(schedule)
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}


func (ss *ScheduleService) DeleteSchedule(id uint) (*model.Schedule, []error) {
	sat, errs := ss.scheduleRepo.DeleteSchedule(id)
	if len(errs) > 0 {
		return nil, errs
	}

	return sat, nil
}

func (ss *ScheduleService) InfosInSchedule(schedule *model.Schedule) ([]model.Info, []error){
	sts, errs := ss.scheduleRepo.InfosInSchedule(schedule)
	if len(errs) > 0 {
		return nil, errs
	}

	return sts, nil
}
