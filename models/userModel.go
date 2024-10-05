package models

import (
    "database/sql"
    "log"
)

type User struct {
    ID       int
    Username string
    Password string
    Role     string
}

// FindUserByUsername finds a user by their username
func FindUserByUsername(db *sql.DB, username string) (*User, error) {
    user := User{}
    
    err := db.QueryRow("SELECT username, password, role FROM users WHERE username = ?", username).
        Scan(&user.Username, &user.Password, &user.Role)
        log.Println(&user)

    if err != nil {
        return nil, err
    }
    return &user, nil
}

// CreateUser creates a new user in the database
func CreateUser(db *sql.DB, username, password, role string) error {
    _, err := db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", username, password, role)
    return err
}
