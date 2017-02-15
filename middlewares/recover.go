package middlewares

import (
	httprouter "github.com/julienschmidt/httprouter"
	"log"
	http "net/http"
)

func Recover(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[FATAL] %+v\n", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next(w, r, ps)
	}
}
