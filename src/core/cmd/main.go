package main

import (
	"helpa/src/core/infra/controllers"
)

func main() {
    router := NewRouter()

    router.Logger.Fatal(router.Start(":80"))
}
