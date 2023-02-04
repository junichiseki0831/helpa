package main

import (
	"fmt"
	"helpa/src/core/infra/controllers"
	"helpa/src/core/domain/userdm"
)

func main() {
	userId := domain.GenerateUserId()
	fmt.Println(userId)
    controllers.CreateRouter().Run(":8080")
}
