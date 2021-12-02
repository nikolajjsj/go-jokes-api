package Config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

var Connection *pgx.Conn

func InitializeDB() {
	newConnection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		Connection = newConnection
	}
}
