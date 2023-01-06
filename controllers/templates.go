package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibePaginaIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"mensagem": "Bem vindo(a) ao site",
	})
}
