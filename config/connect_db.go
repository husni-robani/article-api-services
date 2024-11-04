package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
) 

var DB *sql.DB

func ConnectDatabase(){
	dbConfig := mysql.Config{
		User:   os.Getenv("DB_USER"),
        Passwd: os.Getenv("DB_PASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "article_api_services",
	}

	var err error
	DB, err = sql.Open("mysql", dbConfig.FormatDSN())

	if err != nil {
		log.Fatalf("Error opening database:\n%v\n", err)
	}
	
	if err := DB.Ping(); err != nil{
		log.Fatalf("Error ping database: \n%v\n", err)
	}

	fmt.Println("Database Connected!!")
}