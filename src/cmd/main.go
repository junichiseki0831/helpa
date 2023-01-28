package main

import (
	"helpa/src/core/infra/controllers"
)

func main() {
    controllers.CreateRouter().Run(":8080")
}
