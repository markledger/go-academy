package main

import (
	"api/internal/config"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"log/slog"
	"net/http"
	"os"
)

const portNumber = ":8080"

type customHandler struct {
	slog.Handler
}

func (h *customHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(traceIdKey).(uuid.UUID); ok {
		r.AddAttrs(slog.String(traceIdKey, traceID.String()))
	}
	return h.Handler.Handle(ctx, r)
}

var Ready = make(chan struct{})

var app config.AppConfig

func Init() context.Context {
	baseHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})
	handler := &customHandler{Handler: baseHandler.WithGroup("GoAcademy-API")}
	logger := slog.New(handler)
	slog.SetDefault(logger)

	ctx, _ := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, traceIdKey, uuid.NewString())
	slog.InfoContext(ctx, "Starting")

	return ctx
}

// main is the main function
func main() {
	ctx := Init()
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    portNumber,
		Handler: contextMiddleware(ctx, traceIDMiddleware(routes(mux))),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
