package controllers

import (
	"models"
	"net/http"

	"bancoDados"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	bancoDados.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func BuscaPeloId(c *gin.Context) {
	var result models.Aluno
	id := c.Params.ByName("id")
	bancoDados.DB.First(&result, id)

	if result.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(200, result)
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
