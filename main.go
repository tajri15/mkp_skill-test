package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/tajri15/mkp_skill-test/controllers"
	"github.com/tajri15/mkp_skill-test/middleware"
	"github.com/tajri15/mkp_skill-test/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	models.ConnectDatabase()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "API is running", "time": time.Now()})
	})

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/users")
		{
			auth.POST("/login", controllers.Login)
		}

		schedules := v1.Group("/schedules")
		{
			schedules.GET("/", controllers.GetAllSchedules)
			schedules.GET("/:id", controllers.GetScheduleByID)

			adminRoutes := schedules.Group("/")
			adminRoutes.Use(middleware.AuthMiddleware("admin"))
			{
				adminRoutes.POST("/", controllers.CreateSchedule)
				adminRoutes.PUT("/:id", controllers.UpdateSchedule)
				adminRoutes.DELETE("/:id", controllers.DeleteSchedule)
			}
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
