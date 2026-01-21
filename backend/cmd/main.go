package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/acmota2/musmgr/backend/internal/config"
	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	platform "github.com/acmota2/musmgr/backend/internal/platform/file_access"
	"github.com/acmota2/musmgr/backend/internal/server"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	initConfig, err := config.New()
	if err != nil {
		log.Panic(err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	pool, err := pgxpool.New(ctx, initConfig.DatabaseUrl)
	if err != nil {
		log.Panic(err)
	}
	defer pool.Close()

	log.Println("Successfully connected to the database")

	storageConfig, err := platform.NewStorage(&initConfig.StorageConfig)
	if err != nil {
		log.Panic(err)
	}

	handler := &controller.Handler{
		Pool:    pool,
		Queries: model.New(pool),
		Storage: storageConfig,
	}

	adminRouter := server.NewAdminRouter(&initConfig, handler)
	publicRouter := server.NewPublicRouter(&initConfig, handler)

	go adminRouter.Run(fmt.Sprintf(":%s", initConfig.AdminPort))
	go publicRouter.Run(fmt.Sprintf(":%s", initConfig.PublicPort))
	<-ctx.Done()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")
}
