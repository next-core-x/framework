package tests

import (
	"github.com/gin-gonic/gin"
	"next-core-x/router"
	"testing"
)

func TestNewCustomRouter(t *testing.T) {
	router := router.NewCustomRouter()
	if router == nil {
		t.Error("NewCustomRouter returned nil")
	}
	router.Group("/api/v1").Handle("GET", "/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GET /users"})
	})
	//	For Learning Purpose.....
}
