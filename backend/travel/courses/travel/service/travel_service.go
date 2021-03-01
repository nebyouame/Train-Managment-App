package service


import (
	"github.com/travel/courses/travel"
	"github.com/travel/courses/model"
)

type TravelService struct {
	travelRepo travel.TravelRepository
}

func NewTravelService(TravelRepos travel.TravelRepository) travel.TravelService {

	return &TravelService{travelRepo: TravelRepos}
}


func (cs *TravelService) Travels() ([]model.Travel, []error) {
	cll, errs := cs.travelRepo.Travels()
	if len(errs) > 0 {
		return nil, errs
	}
	return cll, errs
}

func (cs *TravelService) Travel(id uint) (*model.Travel, []error) {
	cmnts, errs := cs.travelRepo.Travel(id)

	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

func (cs *TravelService) StoreTravel(travel *model.Travel) (*model.Travel, []error) {
	cmnts, errs := cs.travelRepo.StoreTravel(travel)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

func (cs *TravelService) UpdateTravel(travel *model.Travel) (*model.Travel, []error) {
	cmnts, errs := cs.travelRepo.UpdateTravel(travel)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

func (cs *TravelService) DeleteTravel(id uint) (*model.Travel, []error) {
	cmnts, errs := cs.travelRepo.DeleteTravel(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

func (cs *TravelService) TravelExists(travelName string) bool {
	exists := cs.travelRepo.TravelExists(travelName)
	return exists
}
