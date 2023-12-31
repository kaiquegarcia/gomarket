package cmd

import (
	"context"
	"fmt"
	route "gomarket/cmd/http"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *application) RunWeb() {
	// Handlers

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	engine := gin.New()
	// TODO: CORS middleware
	// TODO: requestID middleware
	// TODO: logging
	engine = route.SetupRoutes(engine, app.productHandlers)

	srv := &http.Server{
		Addr:    ":8083",
		Handler: engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("HTTP listen error: " + err.Error())
		}
	}()

	<-ctx.Done()

	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic("server forced to shutdown: " + err.Error())
	}

	fmt.Println("shutdown finished")
}
