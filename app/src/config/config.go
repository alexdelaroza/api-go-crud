package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// string de conexao com o MySQL
	StringConexaoBanco = ""
	// Porta onde a API esta rodando
	Porta = 0
	// JwtSecret - chave utilizada para assinar o token
	JwtSecret []byte
)

// Carregar vai inicializar as variaveis de ambiente
func CarregarConfig() {
	var erro error

	erro = godotenv.Load()
	if erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil {
		Porta = 9000
		fmt.Println("ERRO - A Porta foi alterada para:", Porta)
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_BANCO"),
	)
	fmt.Printf("Banco de Dados\n  Usuário: %s\n  Banco..: %s\n", os.Getenv("DB_USUARIO"), os.Getenv("DB_BANCO"))

	JwtSecret = []byte(os.Getenv("JWTSECRET_KEY"))
}
