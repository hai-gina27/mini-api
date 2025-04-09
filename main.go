package main

import (
	"mini-api/config"
	"mini-api/controllers"
	"mini-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Jurusan{}, &models.Mahasiswa{})

	// Routing
	r.GET("/jurusan", controllers.GetJurusan)
	r.POST("/jurusan", controllers.CreateJurusan)
	r.PUT("/jurusan/:id", controllers.UpdateJurusan)
	r.DELETE("/jurusan/:id", controllers.DeleteJurusan)

	r.GET("/mahasiswa", controllers.GetMahasiswa)
	r.POST("/mahasiswa", controllers.CreateMahasiswa)
	r.PUT("/mahasiswa/:nim", controllers.UpdateMahasiswa)
	r.DELETE("/mahasiswa/:nim", controllers.DeleteMahasiswa)

	r.GET("/mahasiswa-jurusan", controllers.GetMahasiswaWithJurusan)

	r.Run(":8083")
}
