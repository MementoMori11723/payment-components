package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v78"
	_ "github.com/stripe/stripe-go/v78"
)

func server() {
	fs := http.FileServer(http.Dir("./pages"))
	http.Handle("/", fs)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{Addr: ":" + port, Handler: nil}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	fmt.Println("Server is running on http://localhost:" + port)
	fmt.Println("Press 'q' and 'Enter' to stop the server")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "q" {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from input: %v", err)
	}
	fmt.Println("Shutting down the server...")
	if err := server.Close(); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Server stopped.")
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found or error loading .env file")
	}
	server()
	stripe.Key = os.Getenv("STRIPE_KEY")
}
