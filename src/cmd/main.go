package main

import (
	"context"
	"fmt"
	infra "helpa/src/core/infra/rdbimpl"
	"log"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
)

func main() {
	db, err := infra.NewDB()
	if err != nil {
		log.Fatalf("エラー: %v", err)
	}
	userRepo := infra.NewUserRepositoryImpl(db)
	users, err := userRepo.FindByName(context.Background(), "testName")
	if err != nil {
		log.Fatalf("Failed to find users by name: %v", err)
	}
	fmt.Println(users)
}
