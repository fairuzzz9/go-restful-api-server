package main

import (
	"context"
	"net/http"
	"time"

	"go-skeleton-rest-app/internal/routes"
	"go-skeleton-rest-app/pkg/utilities"

	"gitlab.com/pos_malaysia/golib/logs"
)

var echoPortNumber = "1234"

// @title Go Rest Skeleton Rest APP
// @version 1.0
// @description This is the Go Skeleton Rest APP

// @contact.name Go Skeleton Rest APP Developer
// @contact.email boojiun@pos.com.my

// @host localhost:1234
// @BasePath /
func main() {

	// setup Echo to use our golib/logs
	e := setupEcho()

	// Initialize and pass the zero logger to the routes.
	routes.InitRoutes(e)

	// Start server by spinning a goroutine so that it will become non-blocking
	go func() {
		if err := e.Start(":" + echoPortNumber); err != nil && err != http.ErrServerClosed {
			logs.Fatal().Msg("shutting down the server")
		}
	}()

	// Gracefully shutdown database, etc services before shutting down Echo server
	// with 10 seconds timeout
	utilities.Graceful(utilities.StopWaitWrapper(func(ctx context.Context) {

		logs.Close()
		e.Shutdown(ctx)
		logs.Info().Msg("Track And Trace API server shutdown")

	}, time.Duration(10)*time.Second))
}
