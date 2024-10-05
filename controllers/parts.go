package controllers

import (
    "net/http"
    "strconv"
    "GoProject/config"
    "GoProject/models"
    "github.com/gin-gonic/gin"
)

func ListParts(c *gin.Context) {
    parts, err := models.GetAllParts(config.DB)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve parts"})
        return
    }
    c.JSON(http.StatusOK, parts)
}

func CreatePart(c *gin.Context) {
    var input struct {
        Name        string  `json:"name"`
        Description string  `json:"description"`
        Price       float64 `json:"price"`
		Stock       int64   `json : "stock"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := models.CreatePart(config.DB, input.Name, input.Description, input.Price,input.Stock)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create part"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Part created successfully"})
}

func GetPart(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid part ID"})
        return
    }

    part, err := models.GetPartByID(config.DB, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Part not found"})
        return
    }

    c.JSON(http.StatusOK, part)
}
