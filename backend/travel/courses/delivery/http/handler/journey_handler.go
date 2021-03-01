package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"

	"github.com/travel/courses/journey"
	"github.com/travel/courses/delivery/http/responses"
	"github.com/travel/courses/model"
	"log"
	"net/http"
	"strconv"
)

type JourneyHandler struct {
	journeyService journey.JourneyService
}

func NewJourneyHander(crService journey.JourneyService) *JourneyHandler {
	return &JourneyHandler{journeyService: crService}
}

func (mh *JourneyHandler) GetJourneys(w http.ResponseWriter, r *http.Request) {

	movies, errs := mh.journeyService.Journeys()

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Can't fetch Journey"))
		return
	}

	responses.JSON(w, http.StatusOK, movies)
	return
}

func (mh *JourneyHandler) PostJourney(w http.ResponseWriter, r *http.Request) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	movie := &model.Journey{}

	err := json.Unmarshal(body, movie)

	if err != nil {
		print("This " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	log.Println(movie.Source)

	movie, errs := mh.journeyService.StoreJourney(movie)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//p := fmt.Sprintf("/v1/admin/comments/%d", comment.ID)
	//w.Header().Set("Location", p)

	responses.JSON(w, http.StatusCreated, movie)
	//w.WriteHeader(http.StatusCreated)
	return
}

func (mh *JourneyHandler) GetJourney(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, err := strconv.Atoi(id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	log.Println(id)

	movie, errs := mh.journeyService.Journey(uint(idd))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(movie, "", "\t\t")
	log.Println(movie)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (mh *JourneyHandler) JourneyUpdate(w http.ResponseWriter, r *http.Request) {
	var u model.Journey
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, err := strconv.Atoi(id)
	//log.Println(idd)
	usr, errs := mh.journeyService.Journey(uint(idd))

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant fetch Course"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	usr, errs = mh.journeyService.UpdateJourney(&u)

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant update Journey"))
		return
	}
	responses.JSON(w, http.StatusOK, &usr)
	return
}

func (mh *JourneyHandler) JourneyDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not given"))
		return
	}
	idd, _ := strconv.Atoi(id)
	u, errs := mh.journeyService.DeleteJourney(uint(idd))
	if len(errs) != 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Couldn't delete journey"))
		return
	}
	responses.JSON(w, http.StatusNoContent, u)
	return

}