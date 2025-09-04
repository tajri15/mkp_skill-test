package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tajri15/mkp_skill-test/models"
)

type ScheduleInput struct {
	MovieID   uint    `json:"movie_id" binding:"required"`
	TheaterID uint    `json:"theater_id" binding:"required"`
	StartTime string  `json:"start_time" binding:"required"`
	EndTime   string  `json:"end_time" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

func CreateSchedule(c *gin.Context) {
	var input ScheduleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startTime, err1 := time.Parse(time.RFC3339, input.StartTime)
	endTime, err2 := time.Parse(time.RFC3339, input.EndTime)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format. Use ISO 8601 (RFC3339), e.g., '2025-09-05T20:00:00+07:00'"})
		return
	}

	schedule := models.Showtime{
		MovieID:   input.MovieID,
		TheaterID: input.TheaterID,
		StartTime: startTime,
		EndTime:   endTime,
		Price:     input.Price,
	}

	if err := models.DB.Create(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		return
	}

	c.JSON(http.StatusCreated, schedule)
}

func GetAllSchedules(c *gin.Context) {
	var schedules []models.Showtime
	models.DB.Find(&schedules)
	c.JSON(http.StatusOK, schedules)
}

func GetScheduleByID(c *gin.Context) {
	var schedule models.Showtime
	if err := models.DB.Where("id = ?", c.Param("id")).First(&schedule).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}
	c.JSON(http.StatusOK, schedule)
}

func UpdateSchedule(c *gin.Context) {
	var schedule models.Showtime
	if err := models.DB.Where("id = ?", c.Param("id")).First(&schedule).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	var input ScheduleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startTime, _ := time.Parse(time.RFC3339, input.StartTime)
	endTime, _ := time.Parse(time.RFC3339, input.EndTime)

	updateData := models.Showtime{
		MovieID:   input.MovieID,
		TheaterID: input.TheaterID,
		StartTime: startTime,
		EndTime:   endTime,
		Price:     input.Price,
	}

	models.DB.Model(&schedule).Updates(updateData)
	c.JSON(http.StatusOK, schedule)
}

func DeleteSchedule(c *gin.Context) {
	var schedule models.Showtime
	if err := models.DB.Where("id = ?", c.Param("id")).First(&schedule).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}
	models.DB.Delete(&schedule)
	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
}