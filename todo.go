package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Todo struct to represent our todo items
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Global database connection variable
var db *sql.DB

// createTable initializes the todos table in the database
func createTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS gotodos (
		id SERIAL PRIMARY KEY,        -- PostgreSQL uses SERIAL for auto-incrementing integers
		title TEXT NOT NULL,
		completed BOOLEAN DEFAULT FALSE -- PostgreSQL uses BOOLEAN and FALSE/TRUE
	);`

	// Execute the SQL statement
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create todos table: %v", err)
	}
	fmt.Println("Todos table checked/created successfully.")
}

// initDB loads environment variables and establishes the database connection
func initDB() {
	// Load environment variables from .env file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Printf("Warning: .env file not loaded. Using environment variables directly. Error: %v", envErr)
		// Don't fatal here, allow using system environment variables
	}

	var err error
	// Open PostgreSQL database connection
	// Ensure this connection string is correct and accessible from Replit
	// It's good practice to get this from an environment variable for security
	db, err = sql.Open("postgres", "postgresql://AggDb_owner:npg_rnFhVj0JHs1Z@ep-plain-lab-a10rbl4g-pooler.ap-southeast-1.aws.neon.tech/AggDb?sslmode=require&channel_binding=require")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")

	// Create the table (or ensure it exists)
	createTable()
}

// app sets up and runs the Gin router
func app() {
	initDB() // Initialize the database connection and create table
	defer func() {
		if db != nil {
			db.Close() // Ensure the database connection is closed when main exits
			fmt.Println("Database connection closed.")
		}
	}()
	router := gin.Default()

	// Basic route for testing
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Go Todo API!"})
	})

	// Todo API routes
	router.GET("/todos", getTodosHandler)
	router.POST("/todos", createTodoHandler)
	router.GET("/todos/:id", getTodoByIdHandler)
	router.PUT("/todos/:id", updateTodoHandler)
	router.DELETE("/todos/:id", deleteTodoHandler)

	// Run the server on port 8080
	// Replit typically exposes port 8080 by default
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	fmt.Println("Server is running on port 8080") // This line might not be reached if Run blocks
}

// getTodosHandler retrieves all todos from the database
func getTodosHandler(c *gin.Context) {
	var todos []Todo
	rows, err := db.Query("SELECT id, title, completed FROM gotodos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to query todos: %v", err)})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to scan todo row: %v", err)})
			return
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil { // Check for errors during iteration
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error during rows iteration: %v", err)})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// getTodoByIdHandler retrieves a single todo by ID
func getTodoByIdHandler(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	// Use $1 for PostgreSQL placeholders
	err := db.QueryRow("SELECT id, title, completed FROM gotodos WHERE id = $1", id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to query todo by ID: %v", err)})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// createTodoHandler creates a new todo
func createTodoHandler(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request payload: %v", err)})
		return
	}

	// Use $1, $2 for PostgreSQL placeholders and RETURNING id to get the new ID
	query := "INSERT INTO gotodos (title, completed) VALUES ($1, $2) RETURNING id"
	err := db.QueryRow(query, todo.Title, todo.Completed).Scan(&todo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to insert todo: %v", err)})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// updateTodoHandler updates an existing todo
func updateTodoHandler(c *gin.Context) {
	id := c.Param("id")
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request payload: %v", err)})
		return
	}

	// Use $1, $2, $3 for PostgreSQL placeholders
	result, err := db.Exec("UPDATE gotodos SET title = $1, completed = $2 WHERE id = $3", todo.Title, todo.Completed, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update todo: %v", err)})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get rows affected: %v", err)})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found or no changes made"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// deleteTodoHandler deletes a todo by ID
func deleteTodoHandler(c *gin.Context) {
	id := c.Param("id")
	// Use $1 for PostgreSQL placeholder
	result, err := db.Exec("DELETE FROM gotodos WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete todo: %v", err)})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get rows affected: %v", err)})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
