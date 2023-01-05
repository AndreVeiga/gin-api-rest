package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndreVeiga/gin-api-rest/controllers"
	"github.com/AndreVeiga/gin-api-rest/database"
	"github.com/AndreVeiga/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

const CPF = "12345678910"

func criaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Nome Aluno Test",
		CPF:  CPF,
		RG:   "123456789",
	}

	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func deletaAlunoMock() {
	var aluno models.Aluno

	database.DB.Delete(&aluno, ID)
}

func SetupRotasTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestSaudacao(t *testing.T) {
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

func TestBuscaTodosAlunos(t *testing.T) {
	database.ConectaComBancoDeDados()
	endpoint := "/alunos"
	criaAlunoMock()
	defer deletaAlunoMock()
	r := SetupRotasTest()
	r.GET(endpoint, controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", endpoint, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestBuscaPorCPF(t *testing.T) {
	endpoint := "/alunos?cpf=" + CPF
	database.ConectaComBancoDeDados()
	criaAlunoMock()
	defer deletaAlunoMock()
	r := SetupRotasTest()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", endpoint, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	// repostaBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.Code)
}
