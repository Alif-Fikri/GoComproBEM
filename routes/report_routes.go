package routes

import (
    "gocomprobem/controllers"
    "github.com/gin-gonic/gin"
)

func ReportRoutes(router *gin.Engine) {
    reportGroup := router.Group("/reports")
    {
        reportGroup.POST("/", controllers.CreateReport)
        reportGroup.GET("/", controllers.GetAllReports)
        reportGroup.GET("/:id", controllers.GetReportByID) 
        reportGroup.GET("/user/:user_id", controllers.GetReportsByUser)
    }
}
