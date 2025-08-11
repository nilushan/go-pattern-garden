# Pattern Garden

Welcome to Pattern Garden, a collection of classic software design patterns implemented in Go. This repository serves as a practical guide and a reference for understanding and applying these patterns in real-world Go applications. It is also a personal project for learning and experimenting with Go.

## Implemented Design Patterns

Here are the design patterns currently available in this collection:

### 1. Options Pattern

The Options pattern provides a clean and extensible way to handle optional parameters in constructors and other functions. It uses functional options to configure a struct, improving readability and maintainability over tangled constructors or zero-value checks.

**Example:**
The `Server` struct in `patterngarden/patterns/options/server.go` is configured using this pattern.

```patterngarden/patterns/options/server.go#L68-73
func Run() {
	server1 := NewServer("api.example.com",
		WithMaxConnections(20),
		WithPort(9000),
		WithTimeout(60*time.Second),
		WithTLS(true),
	)
```

### 2. Repository Pattern

The Repository pattern abstracts the data layer, providing a clean API for the application to interact with data sources without being tied to a specific implementation. This promotes separation of concerns and makes the application more testable and maintainable.

This project includes:
-   A generic `UserRepository` interface.
-   An `InMemoryUserRepository` for testing and development.
-   A `PostgresUserRepository` for production use with a PostgreSQL database.

**Example:**
The `UserService` is initialized with a `UserRepository` implementation. You can easily switch between the in-memory and Postgres repositories.

```patterngarden/cmd/patterngarden/main.go#L36-41
	fmt.Println("--- Running with Postgres Repository ---")
	pgRepo := repository.NewPostgresUserRepository(db)
	userService := repository.NewUserService(pgRepo)

	// fmt.Println("--- Running with In-Memory Repository ---")
	// memRepo := repository.NewInMemoryUserRepository()
	// userService := repository.NewUserService(memRepo)
```

### 3. Factory Pattern

The Factory pattern provides an interface for creating objects in a superclass but lets subclasses alter the type of objects that will be created. The placeholder for this pattern can be found in `patterngarden/patterns/factory/`.


### 4. Pipeline Pattern

The Pipeline pattern is a design pattern that allows you to chain together a series of processing steps, each of which can modify or transform the input data. The placeholder for this pattern can be found in `patterngarden/patterns/pipeline/`.


## Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

-   [Go](https://go.dev/doc/install) (version 1.24 or later)
-   [Docker](https://www.docker.com/products/docker-desktop/) (for running PostgreSQL)
-   A `make` utility

### Setup and Running

The recommended way to run the project is with Docker, which handles the entire environment setup.

1.  **Clone the repository:**
    ```sh
    git clone <repository-url>
    cd patterngarden
    ```

2.  **Start the services:**
    Use the provided `Makefile` to build and start the Go application and the PostgreSQL database in Docker containers.
    ```sh
    make docker-up
    ```
    On the first run, Docker will automatically initialize the database and create the necessary tables using the `db/init/init.sql` script.

3.  **View Logs:**
    To see the output from the running application and database, you can tail the logs.
    ```sh
    make docker-logs
    ```

4.  **Stop the services:**
    To stop the containers and remove the network, use:
    ```sh
    make docker-down
    ```

### Running Locally without Docker

If you prefer to run the Go application directly on your host machine, you can follow these steps:

1.  **Start the database:**
    You still need Docker to run the PostgreSQL database.
    ```sh
    docker-compose -f deploy/docker-compose.yml up -d db
    ```

2.  **Set up environment variables:**
    The application requires a database connection string. You must set it as an environment variable to point to the Dockerized database.
    ```sh
    export DB_CONNECTION_STRING="postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"
    ```

3.  **Run the application:**
    Use the `make run` command or run the `main.go` file directly.
    ```sh
    make run
    ```

## Project Structure

```
patterngarden/
├── cmd/patterngarden/  # Main application entry point
├── db/init/            # SQL database initialization script
├── deploy/             # Docker and deployment configurations
├── patterns/           # Core design pattern implementations
│   ├── options/        # Options pattern
│   ├── repository/     # Repository pattern
│   └── factory/        # Factory pattern (WIP)
├── go.mod              # Go module definitions
├── go.sum              # Go module checksums
└── Makefile            # Make commands for building, running, and managing Docker
```
