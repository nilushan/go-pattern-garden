package main

import (
	"context"
	"fmt"
	"patterngarden/patterns/options"
	"patterngarden/patterns/repository"
)

func runOptions() {
	options.Run()
}

func runRepository() {

	ctx := context.Background()
	fmt.Println("--- Running with In-Memory Repository ---")
	memRepo := repository.NewInMemoryUserRepository()
	userService := repository.NewUserService(memRepo)

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
