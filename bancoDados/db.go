package bancoDados

import (
	"log"
	"models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	conexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(conexao))

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}

	conn.AutoMigrate(&models.Aluno{})

	DB = conn
}
