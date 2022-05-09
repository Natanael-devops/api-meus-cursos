package main

import (
	"github.com/Natanael-devops/api-cursos-gin/database"
	"github.com/Natanael-devops/api-cursos-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
