package main

import (
	"GoTodoWebbox/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("PORT not set in .env")
	}

	log.Println("Server started -> localhost:" + port)
	err := http.ListenAndServe(":"+port, routes.New())

	if err != nil {
		log.Fatal(err)
	}
}
