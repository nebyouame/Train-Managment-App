package journey

import "github.com/travel/courses/model"

type JourneyService interface {
	Journeys() ([]model.Journey, []error)
	Journey(id uint) (*model.Journey, []error)
	UpdateJourney(journey *model.Journey) (*model.Journey, []error)
	DeleteJourney(id uint) (*model.Journey, []error)
	StoreJourney(journey *model.Journey) (*model.Journey, []error)
}
