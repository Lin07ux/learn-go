package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Method(m string) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}
