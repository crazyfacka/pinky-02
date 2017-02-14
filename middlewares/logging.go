package middlewares

import (
	utils "github.com/crazyfacka/pinky-02/utils"
	httprouter "github.com/julienschmidt/httprouter"
	log "log"
	http "net/http"
)

type ResponseLog struct {
	status int
	http.ResponseWriter
}

func (w *ResponseLog) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func Logging(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rl := &ResponseLog{-1, w}
		duration := utils.ExecutionTime(func() {
			next(rl, r, ps)
		})
		log.Printf("[%s] %d %q %v\n", r.Method, rl.status, r.URL.String(), duration)
	}
}
