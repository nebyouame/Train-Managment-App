package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/travel/courses/travel"
	"github.com/travel/courses/delivery/http/responses"
	"github.com/travel/courses/model"
	"log"
	"net/http"
	"strconv"
)

type TravelHandler struct {
	travelService travel.TravelService
}

func NewTravelHandler(CllService travel.TravelService) *TravelHandler {
	return &TravelHandler{travelService: CllService}

}
func (cc *TravelHandler) GetTravels(w http.ResponseWriter, r *http.Request) {
	travels, errs := cc.travelService.Travels()

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Can't fetch Travel"))
		return
	}

	responses.JSON(w, http.StatusOK, travels)
	return

}
func (ach *TravelHandler) PostTravel(w http.ResponseWriter, r *http.Request) {
	var c model.Travel
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant read from request body"))
	}

	log.Println(ach.travelService.TravelExists(c.Name))
	if ach.travelService.TravelExists(c.Name){
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Travel already exists"))
		return
	}

	travel, errs := ach.travelService.StoreTravel(&c)

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant store Travel"))
		return
	}

	log.Println(travel.Name)

	responses.JSON(w, http.StatusCreated, travel)
	return
}

// GetSingleCinema
func (ach *TravelHandler) GetSingleTravel(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, _ := strconv.Atoi(id)

	travel, errs := ach.travelService.Travel(uint(idd))

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant retrieve Travel"))
		return
	}

	responses.JSON(w, http.StatusOK, travel)
	return
}

func (ach *TravelHandler) TravelUpdate(w http.ResponseWriter, r *http.Request) {
	var c model.Travel
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}

	idd, err := strconv.Atoi(id)
	usr, errs := ach.travelService.Travel(uint(idd))

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant fetch Travel"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	usr, errs = ach.travelService.UpdateTravel(&c)

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant update Travel"))
		return
	}
	responses.JSON(w, http.StatusOK, &usr)
	return
}

func (ach *TravelHandler) TravelDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, _ := strconv.Atoi(id)

	_, errs := ach.travelService.DeleteTravel(uint(idd))

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant delete travel"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

