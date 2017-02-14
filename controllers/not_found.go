package controllers

import (
	"fmt"
	utils "github.com/crazyfacka/pinky-02/utils"
	log "log"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	duration := utils.ExecutionTime(func() {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s not found", r.URL.String())
	})
	log.Printf("[%s] 404 %q %v\n", r.Method, r.URL.String(), duration)
}
