package main

import (
	"context"
	"fmt"
	"github.com/acmota2/musmgr/backend/internal/config"
	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	conn, err := pgxpool.New(ctx, initConfig.DatabaseUrl)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to the database")

	handler := &controller.Handler{
		Queries: model.New(conn),
	}

	adminRouter := server.NewAdminRouter(&initConfig, handler)
	publicRouter := server.NewPublicRouter(&initConfig, handler)

	go publicRouter.Run(fmt.Sprintf(":%s", initConfig.PublicPort))
	go adminRouter.Run(fmt.Sprintf(":%s", initConfig.AdminPort))
	<-ctx.Done()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")
}
