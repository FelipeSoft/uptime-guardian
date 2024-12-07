package middleware

import (
	"net/http"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3)

func Limit(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		next(w, r)
	})
}
