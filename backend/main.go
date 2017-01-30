package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joaodias/time-tracker/backend/infrastructure"
	"github.com/joaodias/time-tracker/backend/interfaces"
	"github.com/joaodias/time-tracker/backend/usecases"
	"net/http"
	"os"
)

func main() {
	dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("DBUSER"), os.Getenv("DBKEY"), os.Getenv("DBHOST"), os.Getenv("DBNAME"))
	dbHandler := infrastructure.NewPostgresHandler(dataSource)
	calendarHandler := &infrastructure.GCalendarHandler{
		Logger: &infrastructure.Logger{},
	}

	timeSessionInteractor := &usecases.TimeSessionInteractor{}
	timeSessionInteractor.TimeSessionRepository = &interfaces.DatabaseTimeSessionRepository{
		DatabaseHandler: dbHandler,
		CalendarHandler: calendarHandler,
	}
	timeSessionInteractor.Logger = &infrastructure.Logger{}

	userInteractor := &usecases.UserInteractor{}
	userInteractor.UserRepository = &interfaces.DatabaseUserRepository{
		DatabaseHandler: dbHandler,
	}
	userInteractor.Logger = &infrastructure.Logger{}

	webHandler := &interfaces.WebHandler{}
	webHandler.TimeSessionInteractor = timeSessionInteractor
	webHandler.UserInteractor = userInteractor
	webHandler.Logger = &infrastructure.Logger{}

	router := mux.NewRouter()
	router.HandleFunc("/timesession", webHandler.ListTimeSessions).Methods("GET")
	router.HandleFunc("/timesession", webHandler.NewTimeSession).Methods("POST")
	router.HandleFunc("/user", webHandler.NewUser).Methods("POST")
	router.HandleFunc("/auth", webHandler.SignIn).Methods("GET")
	port := os.Getenv("PORT")
	timeSessionInteractor.Logger.Log("Server running on port " + port)
	http.ListenAndServe(port, router)
}
