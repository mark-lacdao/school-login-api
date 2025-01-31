package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitializeDB() (*pgxpool.Pool, error) {
	connStr := "postgres://username:password@localhost:5432/school"
	dbpool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return dbpool, nil
}
