package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/stripe/stripe-go/v78"
)

func main() {
	godotenv.Load()

	fs := http.FileServer(http.Dir("./pages"))
	http.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
