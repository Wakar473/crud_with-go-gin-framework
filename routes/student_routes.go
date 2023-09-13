// routes/student_routes.go

package routes

import (
    "github.com/gin-gonic/gin"
    controllers  "crud_with_gin/controllers"
)

func SetupStudentRoutes(router *gin.Engine) {
    studentGroup := router.Group("/students")
    {
        studentGroup.POST("/user", controllers.CreateStudent)
        studentGroup.GET("/user", controllers.GetAllStudents)
        studentGroup.GET("/user:id", controllers.GetStudentByID)
    }
}
