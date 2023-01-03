package routes

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest(e *gin.Engine) {
	e.GET("/alunos", controllers.ExibeTodosAlunos)
	e.POST("/alunos", controllers.CriaNovoAluno)
	e.GET("/alunos/:id", controllers.BuscaPeloId)
	e.DELETE("/alunos/:id", controllers.DeletaAluno)
	e.PATCH("/alunos/:id", controllers.EditarAluno)
}
