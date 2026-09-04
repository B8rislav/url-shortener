package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseService struct {
	pool *pgxpool.Pool
}

var (
	dbService DatabaseService
)

func NewPool(ctx context.Context) {
	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create connection: %v", err.Error())
		os.Exit(1)
	}

	if err := dbpool.Ping(ctx); err != nil {
		dbpool.Close()
		log.Fatalf("could not ping database: %v", err)
	}
	dbService.pool = dbpool
}

func SaveUrl(ctx context.Context, shortUrl string, originalUrl string) {
	query := `INSERT INTO urls (shorturl, originalurl) VALUES (@shorturl, @originalurl)`
	args := pgx.NamedArgs{
		"shorturl":    shortUrl,
		"originalurl": originalUrl,
	}
	_, err := dbService.pool.Exec(ctx, query, args)
	if err != nil {
		log.Println("unable to insert row due to: ", err.Error())
	}
}

func GetUrl(ctx context.Context, shortUrl string) (string, error) {
	query := `SELECT originalurl FROM urls WHERE shorturl = $1`
	var foundUrl string
	err := dbService.pool.QueryRow(ctx, query, shortUrl).Scan(&foundUrl)
	if err != nil {
		return "", err
	}
	return foundUrl, nil
}
