package models

type Aluno struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

func ListaTodosAlunos() []Aluno {
	alunos := []Aluno{
		{Id: 1, Nome: "Elton", CPF: "123.456.789-10", RG: "123456789"},
		{Id: 2, Nome: "Tiago", CPF: "109.876.543-21", RG: "987654321"},
	}

	return alunos
}
