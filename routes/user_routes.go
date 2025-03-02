package routes

import (
    "gocomprobem/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
    userGroup := router.Group("/users")
    {
        userGroup.POST("/", controllers.CreateUser) 
        // userGroup.GET("/", controllers.GetAllUsers)
        // userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.GET("/:id/reports", controllers.GetUserReports)
    }
}