package config


import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectDatabase establishes a connection to the MySQL database
func ConnectDatabase() {
   
    // Example DSN: "user:password@tcp(127.0.0.1:3306)/vehicle_parts"
    dsn := os.Getenv("DB_DSN")

    // Open a connection to the database
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Failed to open the database connection!", err)
    }

    // Test the database connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Failed to connect to the database!", err)
    }

    fmt.Println("Successfully connected to the database!")

    // Assign the connection to the global variable
    DB = db
}
