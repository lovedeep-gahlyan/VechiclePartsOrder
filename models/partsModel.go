package models

import "database/sql"

type Part struct {
    ID          int
    Name        string
    Description string
    Price       float64
}

// GetAllParts retrieves all vehicle parts
func GetAllParts(db *sql.DB) ([]Part, error) {
    rows, err := db.Query("SELECT id, name, description, price FROM parts")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var parts []Part
    for rows.Next() {
        var part Part
        if err := rows.Scan(&part.ID, &part.Name, &part.Description, &part.Price); err != nil {
            return nil, err
        }
        parts = append(parts, part)
    }
    return parts, nil
}

// GetPartByID retrieves a part by its ID
func GetPartByID(db *sql.DB, id int) (*Part, error) {
    part := Part{}
    err := db.QueryRow("SELECT id, name, description, price FROM parts WHERE id = ?", id).
        Scan(&part.ID, &part.Name, &part.Description, &part.Price)
    
    if err != nil {
        return nil, err
    }
    return &part, nil
}

// CreatePart creates a new vehicle part
func CreatePart(db *sql.DB, name, description string, price float64,stock int64) error {
    _, err := db.Exec("INSERT INTO parts (name, description, price,stock) VALUES (?, ?, ?,?)", name, description, price,stock)
    return err
}
