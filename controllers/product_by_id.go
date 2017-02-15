package controllers

import (
	"fmt"
	databases "github.com/crazyfacka/pinky-02/databases"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ProductById(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	sqlite := databases.NewProductsDB("./prod.db")
	data, err := sqlite.GetProductById(ps.ByName("pid"))
	if err != nil {
		w.WriteHeader(err.GetError())
		fmt.Fprintf(w, "%s", err.ToJSON())
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", data.ToJSON())
	}
}
