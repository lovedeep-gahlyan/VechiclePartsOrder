package controllers

import (
    "net/http"
    "strconv"
    "GoProject/config"
    "GoProject/models"
    "github.com/gin-gonic/gin"
	"fmt"
    "reflect"
)

// AddToCart adds a part to the user's cart
func AddToCart(c *gin.Context) {
    var input struct {
        UserID int `json:"user_id"`
        PartID int `json:"part_id"`
        //Quantity int `json:"quantity"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := models.AddToCart(config.DB, input.UserID, input.PartID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Part added to cart"})
}

// ViewCart retrieves the user's cart
func ViewCart(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("user_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
   
   
    cartItems, err := models.GetCartItems(config.DB, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart"})
        return
    }
	fmt.Println(reflect.TypeOf(cartItems))
    c.JSON(http.StatusOK, cartItems)
}
