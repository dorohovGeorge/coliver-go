package main

import (
	"context"
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.JSON(makeRequest())
	})
	log.Fatal(app.Listen(":9090"))
}

func makeRequest() string {
	connStr := "postgresql://postgres:4fFzG5313GCQnLCr@94.103.89.23:5432/postgres"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	start := time.Now()
	var name string
	rows, _ := conn.Query(context.Background(), "select name from products")
	rows.Next()

	err = rows.Scan(&name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(name)
	//fmt.Printf("%d\n", name)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	return name
}
