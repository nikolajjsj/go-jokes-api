package Config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Connection *pgx.Conn

func InitializeDB() {
	// Initialize the .env file, to fetch database url
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	Connection, err = pgx.Connect(context.Background(), databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
