package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/travel/courses/rtoken"

	brepo "github.com/travel/courses/book/repository"
	bserv "github.com/travel/courses/book/service"
	jrepo "github.com/travel/courses/journey/repository"
	jserv "github.com/travel/courses/journey/service"
	srepo "github.com/travel/courses/schedule/repository"
	sserv "github.com/travel/courses/schedule/service"
	urepo "github.com/travel/courses/user/repository"
	userv "github.com/travel/courses/user/service"
	"github.com/travel/courses/delivery/http/handler"
	"github.com/travel/courses/model"
	"net/http"
)

func main() {
	db, err := gorm.Open("postgres", "postgres://postgres:password@localhost/movieevent5?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&model.Journey{})
	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Schedule{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
	db.AutoMigrate(&model.Info{})

	token := rtoken.Service{}

	JourneyRepo := jrepo.NewJourneyGormRepo(db)
	Journeysr := jserv.NewJourneyService(JourneyRepo)

	BookingRepo := brepo.NewBookGormRepo(db)
	Bookingsr := bserv.NewBookService(BookingRepo)

	ScheduleRepo := srepo.NewScheduleGormRepo(db)
	Schedulesr := sserv.NewScheduleService(ScheduleRepo)

	UserRepo := urepo.NewUserGormRepo(db)
	usersr := userv.NewUserService(UserRepo)

	roleRepo := urepo.NewRoleGormRepo(db)
	rolesr := userv.NewRoleService(roleRepo)


	mh := handler.NewJourneyHander(Journeysr)
	bh := handler.NewAdminBookHandler(Bookingsr)
	sh := handler.NewAdminScheduleHandler(Schedulesr)
	uah := handler.NewUserHandler(usersr, rolesr, token)





	router := mux.NewRouter()

	router.HandleFunc("/journeys/{id}", mh.GetJourney).Methods("GET")
	router.HandleFunc("/journeys", mh.GetJourneys).Methods("GET")
	router.HandleFunc("/journeys", mh.PostJourney).Methods("POST")
	router.HandleFunc("/journeys/{id}", mh.JourneyUpdate).Methods("PUT")
	router.HandleFunc("/journeys/{id}", mh.JourneyDelete).Methods("DELETE")

	router.HandleFunc("/bookings/{id}", bh.GetBooking).Methods("GET")
	router.HandleFunc("/bookings", bh.GetBookings).Methods("GET")
	router.HandleFunc("/bookings", bh.PostBooking).Methods("POST")
	router.HandleFunc("/bookings/{id}", bh.BookingUpdate).Methods("PUT")
	router.HandleFunc("/bookings/{id}", bh.BookingDelete).Methods("DELETE")

	router.HandleFunc("/schedules/{id}", sh.GetSchedule).Methods("GET")
	router.HandleFunc("/schedules", sh.GetSchedules).Methods("GET")
	router.HandleFunc("/schedules", sh.PostSchedule).Methods("POST")
	router.HandleFunc("/schedules/{id}", sh.ScheduleUpdate).Methods("PUT")
	router.HandleFunc("/schedules/{id}", sh.ScheduleDelete).Methods("DELETE")

	router.HandleFunc("/login", uah.Login).Methods("POST")
	router.HandleFunc("/signup", uah.SignUp).Methods("POST")
	router.HandleFunc("/logout", uah.Authenticated(uah.Logout)).Methods("POST")
	router.HandleFunc("/user/{id}", uah.Authenticated(uah.UserUpdate)).Methods("PUT")
	router.HandleFunc("/user/{id}", uah.Authenticated(uah.UserDelete)).Methods("DELETE")

	http.ListenAndServe(":8181", router)
	//http.ListenAndServe(":8181", nil)

}


