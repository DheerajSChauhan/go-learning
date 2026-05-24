package main

import (
	"context"
	"go-auth/internal/app"
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {
	//root context for the app
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize app: %s", err)
	}
	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Printf("Failed to close app: %s", err)
		}
	}()

	router := httpserver.NewRouter(a)


	//stamdard go type to run the http server
	srv := &http.Server{
		Addr: ":5000",
		Handler: router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Api running on %s", srv.Addr)
	
	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server closed")
		} else {
			log.Fatalf("Server closed with error: %s", err)
		}
	}
}