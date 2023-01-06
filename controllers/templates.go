package controllers

import (
	"net/http"

	"github.com/AndreVeiga/gin-api-rest/database"
	"github.com/AndreVeiga/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func PaginaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
