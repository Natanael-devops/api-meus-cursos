package controllers

import (
	"net/http"

	"github.com/Natanael-devops/api-cursos-gin/database"
	"github.com/Natanael-devops/api-cursos-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosCursos(c *gin.Context) {
	var cursos []models.Curso
	database.DB.Find(&cursos)
	c.JSON(200, cursos)
}

func CriaNovoCurso(c *gin.Context) {
	var curso models.Curso
	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&curso)
	c.JSON(http.StatusOK, curso)
}

func BuscaCursoPorID(c *gin.Context) {
	var curso models.Curso
	id := c.Params.ByName("id")
	database.DB.First(&curso, id)

	if curso.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Curso não encontrado"})
		return
	}

	c.JSON(http.StatusOK, curso)
}

func DeletaCurso(c *gin.Context) {
	var curso models.Curso
	id := c.Params.ByName("id")
	database.DB.Delete(&curso, id)
	c.JSON(http.StatusOK, gin.H{"data": "Curso deleteado com sucesso!"})
}

func EditaCurso(c *gin.Context) {
	var curso models.Curso
	id := c.Params.ByName("id")
	database.DB.First(&curso, id)

	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&curso).UpdateColumns(curso)
	c.JSON(http.StatusOK, curso)
}
