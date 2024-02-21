package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Conexão com MySQL
	StringConexaoDB = ""

	// Porta API
	Porta = 0
)

// Inicializa variáveis de ambiente
func Carregar() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"))
}
