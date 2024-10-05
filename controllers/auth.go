package controllers

import (
    "net/http"
    "GoProject/models"
    "GoProject/utils"
    //"golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
    "GoProject/config"
    "fmt"
)


func Register(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
        Role     string `json:"role"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    PlanePassword, err := utils.HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        fmt.Println(PlanePassword)
        return
    }

    err = models.CreateUser(config.DB, input.Username, input.Password, input.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}


//---------------------------------------------------------------------------------------

// Login controller for both admin and users
func Login(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := models.FindUserByUsername(config.DB, input.Username)
    if err != nil {
        fmt.Println("Inside username block")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if !utils.CheckPassword(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := utils.GenerateJWT(user.Username,user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token,"Message":"User Successfully logged-in"})
}