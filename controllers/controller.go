package controllers

import (
	"models"
	"net/http"

	"bancoDados"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	cpf := c.Query("cpf")
	
	if len(cpf) > 0 {
		BuscaAlunoPeloCPF(c, cpf)
		return
	}

	var alunos []models.Aluno
	bancoDados.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
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

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	bancoDados.DB.Delete(&aluno, id)

	c.JSON(http.StatusNoContent, nil)
}

func EditarAluno(c *gin.Context) {
	var aluno models.Aluno

	id := c.Params.ByName("id")

	bancoDados.DB.First(&aluno, id)

	if aluno.ID != 0 {
		if err := c.ShouldBindJSON(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		bancoDados.DB.Model(&aluno).UpdateColumns(aluno)

		c.JSON(http.StatusOK, aluno)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
		})
	}
}

func BuscaAlunoPeloCPF(c *gin.Context, cpf string) {
	var aluno models.Aluno

	bancoDados.DB.Find(&aluno, "cpf = ?", cpf)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
	} else {
		c.JSON(http.StatusOK, aluno)
	}
}
