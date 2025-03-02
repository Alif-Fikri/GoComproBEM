package main

import (
    "gocomprobem/config"
    "gocomprobem/routes"
    "github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()
    
    config.ConnectDatabase()
    routes.ReportRoutes(router)
    routes.UserRoutes(router)
    
    router.Run(":8080")
}