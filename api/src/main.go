package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/user/goWebRestApp/api/controllers"

	"github.com/rs/cors"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()
	 handler := cors.Default().Handler(r)


	// Get a  DogparkController instance
	uc := controllers.NewDogparkController()

	// Get a  dogpark resource
	r.GET("/dogpark/", uc.GetDogparks)

	// Get a  dogpark resource
	r.GET("/dogpark/:id", uc.GetDogpark)

	r.POST("/dogpark", uc.CreateDogpark)

	r.DELETE("/dogpark/:id", uc.RemoveDogpark)

	// Fire up the server
	http.ListenAndServe("localhost:3000", handler)
}