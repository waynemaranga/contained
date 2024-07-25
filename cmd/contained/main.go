// / Contained is simple demo of a Go program that runs in a container.
// / It provides a simple HTTP server, and a simple health check endpoint.
// / It also exposes a db connection to a Postgres database and a MongoDB database.
package main

// import (
// "context" // for context, i.e for handling signals
// "database/sql" // for sql database i.e Postgres
// "fmt" // for fmt.Println, fmt.Sprintf etc; for string formatting
// "log" // for log.Fatal, log.Printf etc; for logging
// "net/http" // for http server
// "os" // for os.Getenv, os.Exit etc; for environment variables
// "os/signal" // for signal.Notify; for handling signals
// "syscall" // for syscall.SIGINT, syscall.SIGTERM etc; for system calls
// "time" // for time.Sleep, time.Duration etc; for time related operations
// // those are the standard libraries that are used in this program
// )

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	// mongo driver
)

func main() {
	fmt.Println("Hello, Contained!")

	// start the server
	startServer()

	// connect to the databases
	connectToPostgres()
	// connectToMongoDB()

	// from Postgres, we can run queries.
}

// Simple HTTP server
func startServer() {
	// create a new mux; a mux is a HTTP request multiplexer
	mux := http.NewServeMux()
	// add a handler for the /health endpoint; a handler is a function that handles an HTTP request
	// handling an HTTP request means reading the request, processing it, and writing a response
	mux.HandleFunc("/health", healthHandler)
	// create a new http server
	server := &http.Server{
		Addr:    ":8080", // listen on port 8080
		Handler: mux,     // use the mux we created
	}
	// start the server
	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server: %v", err)
	}

	// wait for a shutdown signal
	waitForShutdown(server)
}

// Health check handler; returns "OK" to indicate the service is healthy
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

// Wait for shutdown signal; gracefully shutdown the server
func waitForShutdown(server *http.Server) {
	// create a channel to receive signals
	c := make(chan os.Signal, 1) // a channel is a communication mechanism that allows one goroutine to send a message to another goroutine

	// notify the channel when we receive a SIGINT or SIGTERM signal
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) // signal.Notify registers c to receive notifications of specified signals

	// block until we receive a signal
	<-c // the <- operator is used to receive a message from a channel

	// create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // context.WithTimeout returns a new Context with an associated timeout
	defer cancel()                                                          // defer is used to schedule a function call to be run after the function completes

	// shutdown the server
	log.Println("Shutting down server")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown: %v", err)
	}
}

// Connect to Postgres instance; read the connection details from environment variables
func connectToPostgres() {
	// read the environment variables
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// create a connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// connect to the database
	db, err := sql.Open("postgres", connStr) // sql.Open opens a database specified by its database driver name and a driver-specific data source name

	if err != nil {
		log.Fatalf("postgres: %v", err)
	}

	// ping the database
	if err := db.Ping(); err != nil {
		log.Fatalf("postgres: %v", err)
	}
}

// // Connect to cloud MongoDB instance; read the connection details from environment variables
// func connectToMongoDB() {
// 	// read the environment variables
// 	uri := os.Getenv("MONGODB_URI")

// }

// // Get data from database given in function argument
// func getDataFromDB(db *sql.DB) {
// 	// if Postgres, we can run queries

// 	// query the database
// 	rows, err := db.Query("SELECT * FROM users") // db.Query executes a query that returns rows, typically a SELECT
