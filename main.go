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
	Recover := middlewares.Recover

	router := httprouter.New()
	router.NotFound = controllers.NotFound

	router.GET("/", Logging(Recover(controllers.Alive)))
	router.GET("/products/:pid", Logging(Recover(controllers.ProductById)))

	log.Printf("Serving in :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
