package routes

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest(e *gin.Engine) {
	e.GET("/alunos", controllers.ExibeTodosAlunos)
	e.GET("/alunos/:id", controllers.BuscaPeloId)
	e.POST("/alunos", controllers.CriaNovoAluno)
}
