package main

import (
	"helpa/src/core/infra/controllers"
	"log"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
)

func main() {
	r, err := controllers.CreateRouter()
	if err != nil {
		log.Fatal(err)
	}

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
