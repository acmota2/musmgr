package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/acmota2/musmgr/backend/internal/config"
	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/platform/storage"
	"github.com/acmota2/musmgr/backend/internal/server"
	services "github.com/acmota2/musmgr/backend/internal/services/pdf-generation"

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

	storageConfig, err := storage.NewStorage(&initConfig.StorageConfig)
	if err != nil {
		log.Panic(err)
	}

	baseHandler := controller.NewBaseHandler(
		pool,
		model.New(pool),
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	)

	filesHandler := &controller.FilesHandler{
		Storage:      storageConfig,
		PdfGenerator: services.NewPdfGenerator(services.PDFCPU),
	}

	adminRouter := server.NewAdminRouter(&initConfig, baseHandler, filesHandler)
	publicRouter := server.NewPublicRouter(&initConfig, baseHandler, filesHandler)

	go adminRouter.Run(fmt.Sprintf(":%s", initConfig.AdminPort))
	go publicRouter.Run(fmt.Sprintf(":%s", initConfig.PublicPort))
	<-ctx.Done()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")
}
