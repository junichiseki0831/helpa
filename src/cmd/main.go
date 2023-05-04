package main

import (
	"context"
	_ "embed"
	"encoding/base64"
	app "helpa/src/core/app/userapp"
	"helpa/src/core/domain/shared/vo"
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

	base64String := base64.StdEncoding.EncodeToString(sampleImagePng)

	pass, _ := vo.NewPassword("12345671")
	mail, err := vo.NewEmail("john.doe@example.com")
	img, err := vo.NewImage(base64String)
	req := &app.CreateUserRequest{
		Name:         "John Doe",
		Password:     pass,
		Email:        mail,
		Introduction: "Hello, I'm John!",
		Note:         "Some notes about John",
		Image:        img,
	}

	err = createUserAppService.Exec(context.Background(), req)
	if err != nil {
		panic(err)
	}
}
