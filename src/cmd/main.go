package main

import (
	"context"
	_ "embed"
	app "helpa/src/core/app/userapp"
	infra "helpa/src/core/infra/rdbimpl"
	"log"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
)

//go:embed testdata/sample.png
var sampleImagePng []byte

func main() {
	db, err := infra.NewDB()
	if err != nil {
		log.Fatalf("エラー: %v", err)
	}
	userRepo := infra.NewUserRepositoryImpl(db)
	createUserAppService := app.NewCreateUserAppService(userRepo)
	// users, err := userRepo.FindByName(context.Background(), "testName")
	// if err != nil {
	// 	log.Fatalf("Failed to find users by name: %v", err)
	// }
	// fmt.Println(users)

	req := &app.CreateUserRequest{
		Name:         "John Doe2",
		Password:     "12345672",
		Email:        "john.doe2@example.com",
		Introduction: "Hello, I'm John2!",
		Note:         "Some notes about John2",
		Image:        sampleImagePng,
	}

	err = createUserAppService.Exec(context.Background(), req)
	if err != nil {
		panic(err)
	}
}
