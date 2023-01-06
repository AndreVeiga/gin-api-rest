package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/AndreVeiga/gin-api-rest/controllers"
	"github.com/AndreVeiga/gin-api-rest/database"
	"github.com/AndreVeiga/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

const CPF = "12345678910"
const NOME = "Nome Aluno Test"
const RG = "123456789"

func criaAlunoMock() {
	aluno := models.Aluno{
		Nome: NOME,
		CPF:  CPF,
		RG:   RG,
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

	var aluno models.Aluno
	json.Unmarshal(res.Body.Bytes(), &aluno)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, NOME, aluno.Nome)
	assert.Equal(t, CPF, aluno.CPF)
	assert.Equal(t, RG, aluno.RG)
}

func TestBuscaPorId(t *testing.T) {
	database.ConectaComBancoDeDados()
	criaAlunoMock()
	endpoint := "/alunos/" + strconv.Itoa(ID)
	defer deletaAlunoMock()
	r := SetupRotasTest()
	r.GET("/alunos/:id", controllers.BuscaPeloId)

	req, _ := http.NewRequest("GET", endpoint, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var aluno models.Aluno
	json.Unmarshal(res.Body.Bytes(), &aluno)

	assert.Equal(t, http.StatusOK, res.Code)

	assert.Equal(t, NOME, aluno.Nome)
	assert.Equal(t, CPF, aluno.CPF)
	assert.Equal(t, RG, aluno.RG)
}

func TestDeletaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	criaAlunoMock()
	endpoint := "/alunos/" + strconv.Itoa(ID)
	r := SetupRotasTest()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	req, _ := http.NewRequest("DELETE", endpoint, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNoContent, res.Code)
}

func TestEditaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	criaAlunoMock()
	defer deletaAlunoMock()
	r := SetupRotasTest()
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	nome := "Novo nome pro aluno"
	aluno := models.Aluno{Nome: nome, CPF: CPF, RG: RG}
	var endpoint = "/alunos/" + strconv.Itoa(ID)

	alunoJSON, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(alunoJSON))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var alunoAtualizado models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoAtualizado)

	assert.Equal(t, http.StatusOK, res.Code)

	assert.Equal(t, nome, alunoAtualizado.Nome)
}
