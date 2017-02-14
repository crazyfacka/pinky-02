package main

import (
	controllers "github.com/crazyfacka/pinky-02/controllers"
	middlewares "github.com/crazyfacka/pinky-02/middlewares"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	Logging := middlewares.Logging

	router := httprouter.New()
	router.NotFound = controllers.NotFound

	router.GET("/", Logging(controllers.Alive))
	router.GET("/products/:pid", Logging(controllers.ProductById))

	log.Printf("Serving in :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
