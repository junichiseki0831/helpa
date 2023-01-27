package main

import (
	"helpa/src/core/infra/controllers"
)

func main() {
    router := controllers.NewRouter()

    router.Logger.Fatal(router.Start(":8080"))
}
