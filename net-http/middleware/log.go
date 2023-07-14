package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Logging() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			handler.ServeHTTP(w, r)
		})
	}
}
