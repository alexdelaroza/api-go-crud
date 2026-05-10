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

	JwtSecret = []byte(os.Getenv("JWTSECRET_KEY"))

	StringConexaoBanco = MontaStringConexaoBD()
	fmt.Printf("Banco de Dados\n  Usuário: %s\n  Banco..: %s\n", os.Getenv("DB_USUARIO"), os.Getenv("DB_BANCO"))

}

// Monta a string de conexão formatada para o MySQL
func MontaStringConexaoBD() string {
	// Buscando as variáveis do .env
	usuario := os.Getenv("DB_USUARIO")
	senha := os.Getenv("DB_SENHA")
	banco := os.Getenv("DB_BANCO")
	// A porta interna do container MySQL é "3306"
	porta := os.Getenv("DB_PORTA")

	// Para rodar Docker => o Host: "db"        => "DB_HOST"
	// Para rodar local  => o Host: "127.0.0.1" => "DB_HOST"
	host := os.Getenv("DB_HOST")

	// Formato: usuario:senha@tcp(host:porta)/nome_do_banco?opcoes
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		usuario,
		senha,
		host,
		porta,
		banco,
	)
}
