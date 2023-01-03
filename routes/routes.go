package routes

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest(e *gin.Engine) {
	e.GET("/alunos", controllers.ExibeTodosAlunos)
	e.GET("/:nome", controllers.Saudacao)
}
