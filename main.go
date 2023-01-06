package main

import (
	"github.com/AndreVeiga/gin-api-rest/database"
	"github.com/AndreVeiga/gin-api-rest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConectaComBancoDeDados()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	routes.HandleRequest(r)

	r.Run()
}
