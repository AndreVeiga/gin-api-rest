package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndreVeiga/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRotasTest() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestBuscaTodosAlunos(t *testing.T) {
	nome := "Elton"
	r := SetupRotasTest()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/"+nome, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	respostaEsperada := `{"message":"Olá ` + nome + `, seja bem vindo."}`
	repostaBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, respostaEsperada, string(repostaBody))
}
