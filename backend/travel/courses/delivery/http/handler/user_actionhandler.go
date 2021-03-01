package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/travel/courses/book"

	"github.com/travel/courses/delivery/http/responses"
	"github.com/travel/courses/model"
	"github.com/travel/courses/journey"
	"github.com/travel/courses/rtoken"
	"github.com/travel/courses/schedule"
	"github.com/travel/courses/travel"
	"github.com/travel/courses/user"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserActionHandler struct {
	csrv   travel.TravelService
	ssrv   schedule.ScheduleService
	msrv   journey.JourneyService
	usrv   user.UserService
	tsrv   rtoken.Service
	bsrv   book.BookService
}

func NewUserActionHandler(cs travel.TravelService, ss schedule.ScheduleService, ms journey.JourneyService, u user.UserService, t rtoken.Service, b book.BookService) *UserActionHandler {
	return &UserActionHandler{csrv: cs, ssrv: ss, msrv: ms, tsrv:t, usrv: u, bsrv: b}
}
func (uah *UserActionHandler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	var u model.User
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, err := strconv.Atoi(id)

	usr, errs := uah.usrv.User(uint(idd))

	curid := usr.ID
	pass := usr.Password
	amo := usr.Amount
	rol := usr.RoleID

	log.Println(curid)
	log.Println(pass)
	log.Println(amo)
	log.Println(rol)
	log.Println(usr.FullName)

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant fetch user"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	passnew, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Password Encryption  failed"))
		return
	}

	u.Password = string(passnew)
	u.RoleID = rol
	u.Amount = amo

	usr, errs = uah.usrv.UpdateUser(&u)

	log.Println(usr.ID)
	log.Println(usr.Password)
	log.Println(usr.Amount)
	log.Println(usr.RoleID)
	log.Println(usr.FullName)

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant update user"))
		return
	}
	responses.JSON(w, http.StatusOK, &usr)
	return
}

func (uah *UserActionHandler) UserDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not given"))
		return
	}
	idd, _ := strconv.Atoi(id)
	_, errs := uah.usrv.DeleteUser(uint(idd))
	if len(errs) != 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Couldn't delete user"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

func (uah *UserActionHandler) Journeys(w http.ResponseWriter, r *http.Request){

	journey, errs := uah.msrv.Journeys()
	if len(errs) > 0{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant Fetch Journey"))
	}
	responses.JSON(w, http.StatusOK, journey)
}

func (uah *UserActionHandler) SingleJourney(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, err := strconv.Atoi(id)

	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Can't get Journey ID"))
	}

	curJourney, errs := uah.msrv.Journey(uint(idd))
	if len(errs) > 0{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant Fetch Journey"))
	}
	responses.JSON(w, http.StatusOK, curJourney)
}




func (uah *UserActionHandler) Travels(w http.ResponseWriter, r *http.Request){
	travels, errs := uah.csrv.Travels()
	if len(errs) > 0{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant Fetch Travels"))
	}
	responses.JSON(w, http.StatusOK, travels)
}
func (uah *UserActionHandler) GetSingleCinema(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, _ := strconv.Atoi(id)

	travel, errs := uah.csrv.Travel(uint(idd))

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant retrieve Travel"))
		return
	}

	responses.JSON(w, http.StatusOK, travel)
	return
}

func (uah *UserActionHandler) TravelSchedule(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	travelid, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not passed"))
		return
	}
	id, _ := strconv.Atoi(travelid)


	log.Println(r.URL.Query())
	travelSchedules, errs := uah.ssrv.Schedule(uint(id))
	if len(errs) > 0{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("can't get Travel Schedules"))
	}
	responses.JSON(w, http.StatusOK, travelSchedules)

}

func (uah *UserActionHandler) Bookings(w http.ResponseWriter, r *http.Request){
	_token := r.Header.Get("Authorization")
	_token = strings.Replace(_token, "Bearer ", "", 1)
	fmt.Println(_token)
	claim, err := uah.tsrv.GetClaims(_token)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	var sm []model.Schedule
	user := claim.User
	log.Println(user.FullName)
	books, errs := uah.bsrv.Book(user.ID)
	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	for _, books := range books {
		s, _ := uah.ssrv.Schedule(&books)
		sm = append(sm, *s)
	}
	responses.JSON(w, http.StatusOK, sm)
}
