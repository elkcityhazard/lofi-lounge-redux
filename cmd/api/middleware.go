package main

import "net/http"

//	SetContentType is middleware that ensures the Content-Type is always
//	set to application/json for the rest api

func SetContentType(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
