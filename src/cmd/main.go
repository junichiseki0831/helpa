package main

import (
	"helpa/src/core/infra/controllers"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
)

func main() {
	controllers.CreateRouter().Run(":8080")
}
