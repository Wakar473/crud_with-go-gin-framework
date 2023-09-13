// controllers/subject_controller.go

package controllers

import (
    "context"
    "github.com/gin-gonic/gin"
    // "go.mongodb.org/mongo-driver/bson/primitive" 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "net/http"
    "crud_with_gin/models"
)

var subjectCollection *mongo.Collection

func init() {
    // Initialize the MongoDB collection for subjects.
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, _ := mongo.NewClient(clientOptions)
    ctx := context.TODO()
    _ = client.Connect(ctx)
    subjectCollection = client.Database("class_db").Collection("subjects")
}

func CreateSubject(c *gin.Context) {
    var subject models.Subject
    if err := c.ShouldBindJSON(&subject); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx := context.TODO()
    _, err := subjectCollection.InsertOne(ctx, &subject)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Subject created successfully"})
}

func GetAllSubjects(c *gin.Context) {
    ctx := context.TODO()
    cursor, err := subjectCollection.Find(ctx, nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(ctx)

    var subjects []models.Subject
    if err := cursor.All(ctx, &subjects); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, subjects)
}
