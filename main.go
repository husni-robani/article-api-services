package main

import (
	"github.com/husni-robani/article-api-services/config"
	"github.com/husni-robani/article-api-services/routers"
)

func main() {
	config.ConnectDatabase()	

	router := routers.SetupRouter()
	router.Run(":8080")
}