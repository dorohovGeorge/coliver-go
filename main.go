package main

import (
	"context"
	_ "database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"answer": makeRequest(),
		})
	})
	r.Run()

}

func makeRequest() string {
	connStr := "postgresql://postgres:4fFzG5313GCQnLCr@localhost:5432/postgres"
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
