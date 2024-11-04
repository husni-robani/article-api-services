package main

import (
	"log"

	"github.com/husni-robani/article-api-services/config"
	"github.com/husni-robani/article-api-services/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatalf("Error while load .env file: %v\n", err)
	}
	config.ConnectDatabase()	

	router := routers.SetupRouter()
	router.Run(":8080")
}