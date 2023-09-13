// main.go

// package main

// import (
//     "github.com/gin-gonic/gin"
//     "crud_with_gin/routes"
// )

// func main() {
//     // Initialize Gin
//     router := gin.Default()

//     // Set up routes for subjects and students
//     routes.SetupSubject
// }



// main.go

package main

import (
    "github.com/gin-gonic/gin"
    "crud_with_gin/routes"
)

func main() {
    // Initialize Gin
    router := gin.Default()

    // Set up routes
    routes.SetupStudentRoutes(router)

    // Start the server
    router.Run(":8080")
}
