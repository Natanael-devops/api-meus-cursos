package routes

import (
	"github.com/Natanael-devops/api-cursos-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/cursos", controllers.ExibeTodosCursos)
	r.GET("/cursos/:id", controllers.BuscaCursoPorID)
	r.POST("/cursos", controllers.CriaNovoCurso)
	r.DELETE("/cursos/:id", controllers.DeletaCurso)
	r.PATCH("/cursos/:id", controllers.EditaCurso)
	r.Run(":9000")
}
