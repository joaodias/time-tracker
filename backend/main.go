package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joaodias/time-tracker/backend/infrastructure"
	"github.com/joaodias/time-tracker/backend/interfaces"
	"github.com/joaodias/time-tracker/backend/usecases"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var env *string

// Specify command line flags. env let's you set the environment. Typically you
// can choose production or development. Default to development.
func init() {
	env = flag.String("env", "development", "a string")
}

func main() {
	flag.Parse()
	// This needs to be changed depending on the environment you are working.
	err := godotenv.Load(*env + ".env")
	if err != nil {
		log.Panic("Error loading .env file")
	}
	dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("DBUSER"), os.Getenv("DBKEY"), os.Getenv("DBHOST"), os.Getenv("DBNAME"))
	dbHandler := infrastructure.NewPostgresHandler(dataSource)

	timeSessionInteractor := &usecases.TimeSessionInteractor{}
	timeSessionInteractor.TimeSessionRepository = &interfaces.DatabaseTimeSessionRepository{
		DatabaseHandler: dbHandler,
	}
	timeSessionInteractor.Logger = &infrastructure.Logger{}

	webHandler := &interfaces.WebHandler{}
	webHandler.TimeSessionInteractor = timeSessionInteractor
	webHandler.Logger = &infrastructure.Logger{}

	router := mux.NewRouter()
	router.HandleFunc("/timesession", webHandler.ListTimeSessions).Methods("GET")
	router.HandleFunc("/timesession", webHandler.NewTimeSession).Methods("POST")
	port := os.Getenv("PORT")
	timeSessionInteractor.Logger.Log("Server running on port " + port)
	http.ListenAndServe(port, router)
}
