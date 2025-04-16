package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const traceIdKey = "traceID"

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		traceID := uuid.NewString()
		fmt.Println("contextMiddleware traceId:", traceID)
		requestContext := context.WithValue(r.Context(), traceIdKey, traceID)
		next.ServeHTTP(w, r.WithContext(requestContext))
	})
}

func LogTraceId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("THE TRACEID in log middleware:", r.Context().Value(traceIdKey))
		next.ServeHTTP(w, r)
	})
}
