package models

import (
    "database/sql"
    "fmt"
)

// CartItem represents an item in the user's cart
type CartItem struct {
    id      int `json:"id"`
    user_id   int `json:"user_id"`
    part_id   int `json:"part_id"`
    //Quantity int `json:"quantity"`
}

// AddToCart adds a part to the user's cart
func AddToCart(db *sql.DB, user_id, part_id int) error {
    _, err := db.Exec("INSERT INTO carts (user_id, part_id) VALUES (?, ?)", user_id, part_id)
    return err
}

// GetCartItems retrieves all items in the user's cart
func GetCartItems(db *sql.DB, user_id int) ([]CartItem, error) {
    rows, err := db.Query("SELECT id, user_id,part_id FROM carts WHERE user_id = ?", user_id)
   // fmt.Println("ROWS",&rows)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []CartItem
    for rows.Next() {
        var item CartItem
        if err := rows.Scan(&item.id, &item.user_id, &item.part_id); err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    fmt.Println("ITEMS",items)
    return items, nil
}

// // UpdateCartItem updates the quantity of a cart item
// func UpdateCartItem(db *sql.DB, id, quantity int) error {
//     _, err := db.Exec("UPDATE cart SET quantity = ? WHERE id = ?", quantity, id)
//     return err
// }

// RemoveCartItem removes an item from the cart
func RemoveCartItem(db *sql.DB, id int) error {
    _, err := db.Exec("DELETE FROM carts WHERE id = ?", id)
    return err
}
