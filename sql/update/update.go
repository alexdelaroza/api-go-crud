package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	defer db.Close()

	query := `update usuarios 
	            set   cod_usuario   = ? 
				  ,   nome_usuario  = ? 
				  ,   login_usuario = ?
				  ,   senha_usuario = ? 
				  ,   email_usuario = ? 
				  ,   tipo_usuario  = ?
                where cod_usuario   = ?`

	stmt, _ := db.Prepare(query)

	res, _ := stmt.Exec(1, // cod_usuario
		"Antonio",           // nome_usuario
		"antoniocarlos",     // login_usuario
		"senha123",          // senha_usuario
		"antonio@email.com", // email_usuario
		"admin",             // tipo_usuario
		1,                   // cod_usuario
	)

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Printf("Sucesso! %d linha(s) afetada(s).\n", linhas)

}
