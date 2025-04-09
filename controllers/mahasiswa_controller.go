package controllers

import (
	"mini-api/config"
	"mini-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMahasiswa(c *gin.Context) {
	var mahasiswa []models.Mahasiswa
	config.DB.Find(&mahasiswa)
	c.JSON(http.StatusOK, mahasiswa)
}

func CreateMahasiswa(c *gin.Context) {
	var input models.Mahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func UpdateMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&mahasiswa)
	c.JSON(http.StatusOK, mahasiswa)
}

// Delete Mahasiswa
func DeleteMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
		return
	}

	config.DB.Delete(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa berhasil dihapus"})
}

type MahasiswaJurusanResponse struct {
	Nama    string `json:"nama"`
	NIM     string `json:"nim"`
	Jurusan string `json:"jurusan"`
}

func GetMahasiswaWithJurusan(c *gin.Context) {
	var result []MahasiswaJurusanResponse

	err := config.DB.Table("mahasiswa").
		Select("mahasiswa.nama, mahasiswa.nim, jurusan.jurusan").
		Joins("left join jurusan on mahasiswa.id_jurusan = jurusan.id").
		Scan(&result).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
