package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/stripe/stripe-go/v78"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	// Serve files from the ./pages directory
	fs := http.FileServer(http.Dir("./pages"))
	http.Handle("/", fs)

	// Get the port from environment variables or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create the server
	server := &http.Server{Addr: ":" + port, Handler: nil}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	fmt.Println("Server is running on http://localhost:" + port)
	fmt.Println("Press 'q' and 'Enter' to stop the server")

	// Create a scanner to read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Read input
		input := scanner.Text()
		// Check if input is 'q'
		if input == "q" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from input: %v", err)
	}

	// Attempt to gracefully shut down the server
	fmt.Println("Shutting down the server...")
	if err := server.Close(); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}

	// Give some time for the server to shut down gracefully
	time.Sleep(1 * time.Second)
	fmt.Println("Server stopped.")
}
