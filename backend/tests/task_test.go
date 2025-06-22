package tests

import (
	"bytes"
	"dts-developer-challenge/config"
	"dts-developer-challenge/routes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routes.RegisterTaskRoutes(router)
	return router
}

func TestMain(m *testing.M) {
	config.ConnectDB()
	os.Exit(m.Run())
}

func TestCreateTask(t *testing.T) {
	router := SetupTestRouter()

	task := map[string]interface{}{
		"title":       "Unit Test Task",
		"description": "Created during testing",
		"status":      "pending",
		"due_date":    time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}

	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks/", bytes.NewBuffer(body)) // fixed: no trailing slash in route
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d. Body: %s", w.Code, w.Body.String())
	}
}

func TestGetAllTasks(t *testing.T) {
	router := SetupTestRouter()

	req, _ := http.NewRequest("GET", "/tasks/", nil) // fixed: no trailing slash
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d. Body: %s", w.Code, w.Body.String())
	}
}
