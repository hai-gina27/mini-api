package controllers

import (
	"mini-api/config"
	"mini-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJurusan(c *gin.Context) {
	var jurusan []models.Jurusan
	config.DB.Find(&jurusan)
	c.JSON(http.StatusOK, jurusan)
}

func CreateJurusan(c *gin.Context) {
	var input models.Jurusan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

// Update Jurusan
func UpdateJurusan(c *gin.Context) {
	var jurusan models.Jurusan
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&jurusan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jurusan tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&jurusan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&jurusan)
	c.JSON(http.StatusOK, jurusan)
}

// Delete Jurusan
func DeleteJurusan(c *gin.Context) {
	var jurusan models.Jurusan
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&jurusan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jurusan tidak ditemukan"})
		return
	}

	config.DB.Delete(&jurusan)
	c.JSON(http.StatusOK, gin.H{"message": "Jurusan berhasil dihapus"})
}
