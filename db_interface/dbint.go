package dbinterface

import (
	"context"
	"fmt"
	"os"
	"pricetrack/db"

	"github.com/jackc/pgx/v5"
)

func Createconnection() *pgx.Conn{
    url := "postgresql://selva:pwd@127.0.0.1:26257/scrape?sslmode=disable"
    conn, err := pgx.Connect(context.Background(), url)
    if err != nil {
            fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
    }
    return conn
}

func ListURLS(conn *pgx.Conn) []string{

    defer conn.Close(context.Background())
    q := db.New(conn)

    products, err := q.ListProductURLS(context.Background())
    if err != nil {
            fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
            os.Exit(1)
    }

    return products
}