// controllers/student_controller.go

package controllers

import (
	"context"
	"crud_with_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var studentCollection *mongo.Collection

func init() {
    // Initialize the MongoDB collection for students.
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, _ := mongo.NewClient(clientOptions)
    ctx := context.TODO()
    _ = client.Connect(ctx)
    studentCollection = client.Database("class_db").Collection("student")
}

func CreateStudent(c *gin.Context) {
    var student models.Student
    if err := c.ShouldBindJSON(&student); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx := context.TODO()
    _, err := studentCollection.InsertOne(ctx, &student)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Student created successfully"})
}

func GetAllStudents(c *gin.Context) {
    ctx := context.TODO()
    cursor, err := studentCollection.Find(ctx, nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(ctx)

    var students []models.Student
    if err := cursor.All(ctx, &students); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, students)
}

func GetStudentByID(c *gin.Context) {
    studentID := c.Param("id")

    ctx := context.TODO()
    objID, err := primitive.ObjectIDFromHex(studentID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
        return
    }

    var student models.Student
    err = studentCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&student)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
        return
    }

    c.JSON(http.StatusOK, student)
}
