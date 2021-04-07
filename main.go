package main

import (
	"log"

	"github.com/go_postgres/helper"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		log.Println(err)
	}

	log.Println("Database tables are successfully initialized.")
}
