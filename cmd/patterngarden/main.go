package main

import (
	"patterngarden/patterns/options"
	"patterngarden/patterns/repository"

	// Import the Postgres driver
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	options.Run()
	repository.Run()
}
