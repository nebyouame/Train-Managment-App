package travel

import "github.com/travel/courses/model"
type TravelRepository interface {
	Travels() ([]model.Travel, []error)
	Travel(id uint) (*model.Travel, []error)
	StoreTravel(travel *model.Travel) (*model.Travel, []error)
	UpdateTravel(cinema *model.Travel) (*model.Travel, []error)
	DeleteTravel(id uint) (*model.Travel, []error)
	TravelExists(travelName string) bool
}
