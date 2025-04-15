package main

import (
	"net/http"
)

const traceIdKey = "traceID"

func ContextWithTraceID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {

		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
