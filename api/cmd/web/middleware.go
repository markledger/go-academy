package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const traceIdKey = "traceID"

func contextMiddleware(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx, cancel := context.WithCancel(r.Context())
		context.AfterFunc(ctx, cancel)
		next.ServeHTTP(w, r.WithContext(rctx))
	})
}

func traceIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.NewString()

		traceIDHeader := r.Header[traceIdKey]
		fmt.Println(traceIDHeader)
		if traceIDHeader != nil && len(traceIDHeader) != 0 {
			traceID = traceIDHeader[0]
		}
		w.Header().Add(traceIdKey, traceID)
		r.WithContext(context.WithValue(r.Context(), traceIdKey, traceID))
		next.ServeHTTP(w, r)
	})
}
