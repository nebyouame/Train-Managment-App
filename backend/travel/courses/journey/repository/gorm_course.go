package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/travel/courses/journey"
	"github.com/travel/courses/model"

)

type JourneyGormRepo struct {
	conn *gorm.DB
}

func NewJourneyGormRepo(db *gorm.DB) journey.JourneyRepository{
	return &JourneyGormRepo{conn: db}
}

func (journeyrepo *JourneyGormRepo) Journeys() ([]model.Journey, []error) {
	crs := []model.Journey{}
	errs := journeyrepo.conn.Find(&crs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return crs, errs
}

func (journeyrepo *JourneyGormRepo) Journey(id uint) (*model.Journey, []error) {
	ctg := model.Journey{}
	errs := journeyrepo.conn.First(&ctg, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ctg, errs
}

func (journeyrepo *JourneyGormRepo) UpdateJourney(journey *model.Journey) (*model.Journey, []error) {
	eve := journey
	errs := journeyrepo.conn.Save(eve).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return eve, errs
}

func (journeyrepo *JourneyGormRepo) DeleteJourney(id uint) (*model.Journey, []error) {
	eve, errs := journeyrepo.Journey(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = journeyrepo.conn.Delete(eve, eve.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return eve, errs
}

func (journeyrepo *JourneyGormRepo) StoreJourney(journey *model.Journey) (*model.Journey, []error) {
	mv := journey
	errs := journeyrepo.conn.Create(mv).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return mv, errs
}