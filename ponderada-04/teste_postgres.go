package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func insertUser(db *sql.DB, name string, age int) error {
	_, err := db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", name, age)
	if err != nil {
		return err
	}
	fmt.Print("User inserted successfully!\n")
	return nil
}

func displayUsers(db *sql.DB) error {
    rows, err := db.Query("SELECT name, age FROM users")
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var name string
        var age int
        if err := rows.Scan(&name, &age); err != nil {
            return err
        }
        fmt.Printf("Name: %s, Age: %d\n", name, age)
    }
    if err := rows.Err(); err != nil {
        return err
    }
    return nil
}


func main() {
	// Replace these environment variables with your own PostgreSQL connection details
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "mysecretpassword"
	dbname := "mydatabase"
	sslmode := "disable" // Adjust based on your security requirements

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the PostgreSQL database!")

	// Create a simple table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			age INT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'users' created or already exists!")

	// Now you can perform database operations using the db object
	// For example, you can execute queries or use the database within your application logic.

	// Insert a new user
    if err := insertUser(db, "John", 30); err != nil {
        log.Fatal(err)
    }

    // Retrieve and display users
    if err := displayUsers(db); err != nil {
        log.Fatal(err)
    }
	
}
