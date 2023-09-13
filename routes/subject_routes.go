// routes/subject_routes.go

package routes

import (
    "github.com/gin-gonic/gin"
    "crud_with_gin/controllers"
)

func SetupSubjectRoutes(router *gin.Engine) {
    subjectGroup := router.Group("/subjects")
    {
        subjectGroup.POST("/user", controllers.CreateSubject)
        subjectGroup.GET("/user", controllers.GetAllSubjects)
    }
}
