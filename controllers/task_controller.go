package controllers

import (
	"context"
	"dts-developer-challenge/config"
	"dts-developer-challenge/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = primitive.NewObjectID().Hex()
	_, err := config.TaskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create task - %s ", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	err := config.TaskCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func GetAllTasks(c *gin.Context) {
	cursor, err := config.TaskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task
	for cursor.Next(context.TODO()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

func UpdateTaskStatus(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.TaskCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": body.Status}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := config.TaskCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
