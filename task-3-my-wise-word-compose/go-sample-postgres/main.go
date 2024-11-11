package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {

	// Environment variables for database configuration

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Database connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open database connection (does not connect yet)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Create Database Connection failed:", err)
		return
	}

	fmt.Printf("DB_USER: %s\n", dbUser)
	fmt.Printf("DB_PASS: %s\n", dbPass)
	fmt.Printf("DB_NAME: %s\n", dbName)
	fmt.Printf("DB_HOST: %s\n", dbHost)
	fmt.Printf("DB_PORT: %s\n", dbPort)

	// Test database connection
	fmt.Println("Connect to:", connStr)
	if err := db.Ping(); err != nil {
		fmt.Printf("Connection to database failed (DB_HOST: %s): %s\n", dbHost, err)
		return
	}
	fmt.Println("Successfully connected to the database")

	// Initialize Echo
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route for the root path
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Someone hit me!")
		return c.String(http.StatusOK, "The only thing we never get enough of is love; and the only thing we never give enough of is love.")
	})

	// Route for the /connect path
	e.GET("/connect", func(c echo.Context) error {
		// Open a new connection
		conn, err := sql.Open("postgres", connStr)
		if err != nil {
			fmt.Println("Failed to connect to db:", err)
			return c.HTML(http.StatusInternalServerError, "Failed")
		}
		defer conn.Close()

		// Insert access log
		username := "irsyaan" // Replace with your actual username
		query := fmt.Sprintf("INSERT INTO public.%s_access_log (timestamp) VALUES ($1)", username)
		_, err = conn.Exec(query, time.Now())
		if err != nil {
			fmt.Println("Failed to insert log:", err)
			return c.HTML(http.StatusInternalServerError, "Failed")
		}

		fmt.Println("Success connect to db")
		return c.HTML(http.StatusOK, "Success")
	})

	// Start server on specified port
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "80"
	}
	e.Logger.Fatal(e.Start(":" + httpPort))
}
