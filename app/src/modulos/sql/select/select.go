package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	cod_usuario   int
	nome_usuario  string
	login_usuario string
	senha_usuario string
	email_usuario string
	tipo_usuario  string
	data_ult_atu  time.Time
}

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo?parseTime=true")
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer db.Close()

	rows, _ := db.Query("select * from usuarios where cod_usuario = ?", 1)
	defer db.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.cod_usuario, &u.nome_usuario, &u.login_usuario, &u.senha_usuario, &u.email_usuario, &u.tipo_usuario, &u.data_ult_atu)
		fmt.Printf("Usuário: %d - %s (Atualizado em: %s)\n",
			u.cod_usuario, u.nome_usuario, u.data_ult_atu.Format("02/01/2006 15:04:05"))
	}

}
