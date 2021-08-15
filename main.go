package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Ivnkara/ocr-test/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	http.HandleFunc("/from-file", server.FromFile)
	http.HandleFunc("/from-url", server.FromUrl)

	fmt.Println("Server running...")
	err = http.ListenAndServe(os.Getenv("PORT"), nil)
	if err != nil {
		fmt.Println("Error running server")
		os.Exit(2)
	}
}
