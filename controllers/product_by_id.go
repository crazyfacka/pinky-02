package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ProductById(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", ps.ByName("pid"))
}
