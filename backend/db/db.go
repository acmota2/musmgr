package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PsqlDB *pgxpool.Pool

func Connect() *pgxpool.Pool {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	var err error
	PsqlDB, err = pgxpool.New(context.Background(), url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database")
	return PsqlDB
}
