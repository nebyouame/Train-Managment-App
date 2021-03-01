package service

import (
	"github.com/travel/courses/journey"
	"github.com/travel/courses/model"
)

type JourneyService struct {
	journeyRepo journey.JourneyRepository
}

func NewJourneyService(crepo journey.JourneyRepository) journey.JourneyService {
	return &JourneyService{journeyRepo: crepo}
}

func (c *JourneyService) Journeys() ([]model.Journey, []error) {
	mvs, errs := c.journeyRepo.Journeys()
	if len(errs) > 0 {
		return nil, errs
	}
	return mvs, errs
}

func (c *JourneyService) Journey(id uint) (*model.Journey, []error) {
	mov, errs := c.journeyRepo.Journey(id)

	if len(errs) > 0 {
		return mov, errs
	}

	return mov, nil
}

func (c *JourneyService) UpdateJourney(journey *model.Journey) (*model.Journey, []error) {
	eve, errs := c.journeyRepo.UpdateJourney(journey)

	if len(errs) > 0 {
		return nil, errs
	}

	return eve, nil
}

func (c *JourneyService) DeleteJourney(id uint) (*model.Journey, []error) {
	eve, errs := c.journeyRepo.DeleteJourney(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return eve, nil
}

func (c *JourneyService) StoreJourney(journey *model.Journey) (*model.Journey, []error) {
	mvs, errs := c.journeyRepo.StoreJourney(journey)
	if len(errs) > 0 {
		return nil, errs
	}
	return mvs, errs
}