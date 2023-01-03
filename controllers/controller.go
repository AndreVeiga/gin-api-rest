package controllers

import (
	"models"

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
