package main

import (
	"apelsin/api/docs"
	"apelsin/config"
	"apelsin/pkg/cors"
	"apelsin/pkg/logger"
	"apelsin/rest"
	"apelsin/storage/sqlstorage"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// @Title Apelsin
// @description Apelsin crud
// @contact.name API Support
// @Schemes http
// @contact.url https://www.instagram.com/
// @contact.email jasuravazov4@gmail.com
// @license.name Jasur
// @license.url https://www.instagram.com/
func main() {
	// Catch Ctrl^C for graceful shutdown
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt, syscall.SIGTERM) // Non-blocking

	// *** Load config and logger
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "dbo_some_service")

	// Set Host for swagger. In case we need to add swagger to WSO2API Manager
	docs.SwaggerInfo.Host = cfg.HTTPHost + cfg.HTTPPort
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	// *** Storage Initialization
	log.Info("Connecting to db...")
	db, err := sqlx.Connect("postgres", cfg.PostgresURL())
	if err != nil {
		panic(err)
	}
	storage := sqlstorage.New(db, log)
	log.Info("Connected to db...")

	// *** Router Initialization
	r := gin.New()
	r.Use(cors.CORSMiddleware())
	r.Use(gin.Logger(), gin.Recovery())

	// *** REST Initialization
	handler := rest.NewAPI(cfg, log, r, storage)

	srv := &http.Server{
		Addr:    cfg.HTTPPort,
		Handler: handler,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(fmt.Sprintf("Failed To Start REST Server: %s\n", err.Error()))
		}
	}()
	log.Info("REST Server started at port" + cfg.HTTPPort)

	// Wait for Ctrl^C
	OSCall := <-quitSignal

	// Graceful HTTP server shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Info(fmt.Sprintf("\nSystem Call:%+v", OSCall))
	log.Info("GRPC Server Gracefully Shut Down")
	fmt.Printf("system call:%+v", OSCall)

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("REST Server Graceful Shutdown Failed: %s\n", err))
	}
	log.Info("REST Server Gracefully Shut Down")
	cancel()
}
