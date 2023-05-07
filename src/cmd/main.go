package main

import (
	_ "embed"
	"helpa/src/core/infra/controllers"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
)

//go:embed testdata/sample.png
var sampleImagePng []byte

func main() {
	controllers.CreateRouter().Run(":8080")
}
