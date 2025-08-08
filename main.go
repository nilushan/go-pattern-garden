package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"patterngarden/patterns/options"
	"patterngarden/patterns/repository"

	// Import the Postgres driver
	_ "github.com/jackc/pgx/v5/stdlib"
)

func runOptions() {
	options.Run()
}

func runRepository() {

	ctx := context.Background()

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable is not set")
	}

	db, err := sql.Open("pgx", dbConnectionString)
	if err != nil {
		fmt.Printf("Error opening database: %v", err)
		return
	}
	defer db.Close()

	fmt.Println("--- Running with Postgres Repository ---")
	pgRepo := repository.NewPostgresUserRepository(db)
	userService := repository.NewUserService(pgRepo)

	// fmt.Println("--- Running with In-Memory Repository ---")
	// memRepo := repository.NewInMemoryUserRepository()
	// userService := repository.NewUserService(memRepo)

	fmt.Println("Registering user 'John Doe'...")
	user, err := userService.RegisterUser(ctx, "user@example.com", "John Doe")
	if err != nil {
		fmt.Printf("Error registering user: %v", err)
		return
	}
	fmt.Printf("User registered: %v\n", user)

	userByEmail, err := userService.Repo.GetByEmail(ctx, user.Email)
	if err != nil {
		fmt.Printf("Error getting user by email: %v", err)
		return
	}
	fmt.Printf("User found by email: %v\n", userByEmail)

	userById, err := userService.Repo.GetById(ctx, userByEmail.ID)
	if err != nil {
		fmt.Printf("Error getting user by Id: %v", err)
		return
	}
	fmt.Printf("User found by Id: %v\n", userById)
}

func main() {
	runOptions()
	runRepository()

}
