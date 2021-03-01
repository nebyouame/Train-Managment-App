package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/travel/courses/model"
	"github.com/travel/courses/schedule"
	"github.com/travel/courses/delivery/http/responses"
	"log"
	"net/http"
	"strconv"
)

type AdminBookHandler struct {

	bookSrv schedule.BookService
}

func NewAdminBookHandler(bs schedule.BookService) *AdminBookHandler {
	return &AdminBookHandler{bookSrv:bs}
}

func (mh *AdminBookHandler) GetBookings(w http.ResponseWriter, r *http.Request) {

	movies, errs := mh.bookSrv.Books()

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Can't fetch Booking"))
		return
	}

	responses.JSON(w, http.StatusOK, movies)
	return
}

func (mh *AdminBookHandler) PostBooking(w http.ResponseWriter, r *http.Request) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	movie := &model.Book{}

	err := json.Unmarshal(body, movie)

	if err != nil {
		print("This " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	log.Println(movie.ID)

	movie, errs := mh.bookSrv.StoreBook(movie)

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

func (mh *AdminBookHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
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

	movie, errs := mh.bookSrv.Book(uint(idd))

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

func (mh *AdminBookHandler) BookingUpdate(w http.ResponseWriter, r *http.Request) {
	var u model.Book
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, err := strconv.Atoi(id)
	//log.Println(idd)
	usr, errs := mh.bookSrv.Book(uint(idd))

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant fetch Booking"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	usr, errs = mh.bookSrv.UpdateBook(&u)

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant update Booking"))
		return
	}
	responses.JSON(w, http.StatusOK, &usr)
	return
}

func (mh *AdminBookHandler) BookingDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not given"))
		return
	}
	idd, _ := strconv.Atoi(id)
	u, errs := mh.bookSrv.DeleteBook(uint(idd))
	if len(errs) != 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Couldn't delete Booking"))
		return
	}
	responses.JSON(w, http.StatusNoContent, u)
	return

}