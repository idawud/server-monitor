package main

import (
	"context"
	"github.com/idawud/server-monitor/handler"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "server-monitor ", log.LstdFlags)
	ep := handler.NewWebSocketEndpoint(l)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/feedback", ep.MainEndpoint)

	server := &http.Server{
		Addr: ":8080",
		Handler: sm,
		IdleTimeout:120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout:1*time.Second,
	}

	log.Println("Server running on http://localhost:8080/ started at: ", time.Stamp )
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Graceful shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = server.Shutdown(ctx)
}
