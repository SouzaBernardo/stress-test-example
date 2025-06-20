package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/SouzaBernardo/sample-api/internal/routes"
)

func main() {

	mux := http.NewServeMux()
	routes.Metrics(mux)
	routes.Payment(mux)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			return
		}
	}()
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
