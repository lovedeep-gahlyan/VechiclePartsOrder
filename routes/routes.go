package routes

import (
    "github.com/gin-gonic/gin"
    "GoProject/controllers"
    //"GoProject/middleware"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Auth routes
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    // Parts routes
    r.GET("/parts", controllers.ListParts)
    r.GET("/parts/:id", controllers.GetPart)
    r.POST("/part", controllers.CreatePart) 

    r.POST("/addtoCart",controllers.AddToCart)
    r.GET("/getCartItems/:user_id",controllers.ViewCart)

    return r
}
