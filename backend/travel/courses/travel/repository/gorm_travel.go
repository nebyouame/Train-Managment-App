package repository


import (
	"github.com/jinzhu/gorm"
	"github.com/travel/courses/travel"
	"github.com/travel/courses/model"
)


type TravelGormRepo struct {
	conn *gorm.DB
}


func NewTravelGormRepo(db *gorm.DB) travel.TravelRepository {
	return &TravelGormRepo{conn: db}
}
func (cllRepo *TravelGormRepo) Travels() ([]model.Travel, []error) {
	cll := []model.Travel{}
	errs := cllRepo.conn.Find(&cll).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cll, errs
}

//Cinema retrieves a cinema from the database by its id
func (cllRepo *TravelGormRepo) Travel(id uint) (*model.Travel, []error) {
	cll := model.Travel{}
	errs := cllRepo.conn.First(&cll, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cll, errs
}


func (cllRepo *TravelGormRepo) StoreTravel(travel *model.Travel) (*model.Travel, []error) {
	cll := travel
	errs := cllRepo.conn.Create(cll).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cll, errs
}

func (cliRepo *TravelGormRepo) UpdateTravel(travel *model.Travel) (*model.Travel, []error) {
	cli := travel
	errs := cliRepo.conn.Save(travel).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cli, errs
}
func (cliRepo *TravelGormRepo) DeleteTravel(id uint) (*model.Travel, []error) {
	cli, errs := cliRepo.Travel(id)
	if len(errs) > 0 {
		return nil, errs
	}

	errs = cliRepo.conn.Delete(cli, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return cli, errs
}

func (cliRepo *TravelGormRepo) TravelExists(travel string) bool {
	cli := model.Travel{}
	errs := cliRepo.conn.Find(&cli, "name=?", travel).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}