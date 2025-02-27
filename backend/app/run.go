package app

import (
	"context"
	"fmt"
	"gp/backend/api/calls"
	"gp/backend/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Run the server with a port or IP address as an argument
func Run(addr string) {
	log.Println("Starting server")

	err := calls.FetchAll()
	if err != nil {
		log.Fatalf("Error initializing artists: %v", err)
	}

	ctx, cancRefresh := context.WithCancel(context.Background())
	go calls.RefreshDB(ctx)

	gpHandler := handlers.Handlers()
	server := &http.Server{Addr: addr, Handler: gpHandler}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe() failed: %v", err)
		}
	}()
	log.Printf("Server running on %s\n", addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	cancRefresh()

	fmt.Println()
	log.Println("Shutting down server...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server Shutdown Failed: %v", err)
	}
	log.Println("Server stopped")
}
