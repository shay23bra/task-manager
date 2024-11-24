package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"task-manager/models"
	"task-manager/routes"
	"task-manager/utils"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB
var router *gin.Engine

func setup() {
	db = utils.InitDB()
	db.Logger = db.Logger.LogMode(0) // Disable GORM logging during tests
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	routes.RegisterTaskRoutes(router, db)
}

func TestCreateTask(t *testing.T) {
	setup()

	task := models.Task{ // Mock input
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "Pending",
	}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code) // Validate response

	var createdTask models.Task
	json.Unmarshal(resp.Body.Bytes(), &createdTask)

	assert.Equal(t, task.Title, createdTask.Title)
	assert.Equal(t, task.Description, createdTask.Description)
	assert.Equal(t, task.Status, createdTask.Status)
	assert.NotZero(t, createdTask.ID)
}

func TestGetTaskByID(t *testing.T) {
	setup()

	task := models.Task{ // Mock input
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "Pending",
	}
	result := db.Create(&task)
	t.Logf("ID: %v", task.ID)
	assert.NoError(t, result.Error)
	assert.NotZero(t, task.ID, "Task ID should not be zero after creation")
	assert.NotZero(t, result.RowsAffected, "RowsAffected should be greater than zero")

	req, _ := http.NewRequest("GET", "/tasks/"+strconv.Itoa(int(task.ID)), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code) // Validate response

	var retrievedTask models.Task
	json.Unmarshal(resp.Body.Bytes(), &retrievedTask)

	assert.Equal(t, task.Title, retrievedTask.Title)
	assert.Equal(t, task.Description, retrievedTask.Description)
	assert.Equal(t, task.Status, retrievedTask.Status)
}

func TestUpdateTask(t *testing.T) {
	setup()

	task := models.Task{ // Mock input
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "Pending",
	}
	db.Create(&task)

	update := models.Task{ // Mock update input
		Title:       "Updated Task",
		Description: "Updated description",
		Status:      "Completed",
	}
	updateJSON, _ := json.Marshal(update)

	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(int(task.ID)), bytes.NewBuffer(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code) // Validate response

	var updatedTask models.Task
	json.Unmarshal(resp.Body.Bytes(), &updatedTask)

	assert.Equal(t, update.Title, updatedTask.Title)
	assert.Equal(t, update.Description, updatedTask.Description)
	assert.Equal(t, update.Status, updatedTask.Status)
}

func TestDeleteTask(t *testing.T) {
	setup()

	task := models.Task{ // Mock input
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "Pending",
	}
	db.Create(&task)

	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(int(task.ID)), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code) // Validate response

	var deletedTask models.Task
	err := db.First(&deletedTask, task.ID).Error
	assert.Equal(t, gorm.ErrRecordNotFound, err) // Ensure task no longer exists
}
