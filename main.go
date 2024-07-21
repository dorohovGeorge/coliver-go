package main

import (
	"context"
	_ "database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func main() {
	connStr := "postgresql://postgres:4fFzG5313GCQnLCr@localhost:5432/postgres"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	start := time.Now()
	var name int64
	rows, _ := conn.Query(context.Background(), "select id from products")

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//fmt.Printf("%d\n", name)
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

}
