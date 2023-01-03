package main

import (
	"bancoDados"
	"routes"

	"github.com/gin-gonic/gin"
)

func main() {
	bancoDados.ConectaComBancoDeDados()

	r := gin.Default()
	routes.HandleRequest(r)

	r.Run()
}
