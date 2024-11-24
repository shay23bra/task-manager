package routes

import (
	"task-manager/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTaskRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Task Manager API!",
		})
	})

	r.POST("/tasks", func(c *gin.Context) {
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		task.CreatedAt = time.Now()
		db.Create(&task)
		c.JSON(200, task)
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		var task models.Task
		if err := db.First(&task, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(200, task)
	})

	r.GET("/tasks", func(c *gin.Context) {
		var tasks []models.Task
		if err := db.Find(&tasks).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch tasks"})
			return
		}
		c.JSON(200, tasks)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
		var task models.Task
		if err := db.First(&task, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Task not found"})
			return
		}
		c.ShouldBindJSON(&task)
		db.Save(&task)
		c.JSON(200, task)
	})

	r.PUT("/tasks/:id/status", func(c *gin.Context) {
		var task models.Task

		if err := db.First(&task, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Task not found"})
			return
		}

		var input struct {
			Status string `json:"status"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		validStatuses := map[string]bool{"Pending": true, "Done": true, "Hold": true}
		if !validStatuses[input.Status] {
			c.JSON(400, gin.H{"error": "Invalid status"})
			return
		}

		task.Status = input.Status
		if err := db.Save(&task).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update task"})
			return
		}

		c.JSON(200, task)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		db.Delete(&models.Task{}, c.Param("id"))
		c.JSON(200, gin.H{"status": "Deleted"})
	})
}
