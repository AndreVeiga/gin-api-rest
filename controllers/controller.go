package controllers

import (
	"models"
	"net/http"

	"bancoDados"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.ListaTodosAlunos())
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	message := "Olá, tudo bem " + nome

	c.JSON(200, gin.H{
		"message": message,
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	bancoDados.DB.Create(&aluno)

	c.JSON(http.StatusCreated, aluno)
}

