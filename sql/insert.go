package main

import (
	"database/sql"
	"github.com/jbarham/gopgsqldriver"
)

func main(
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nill{
		panic(err)
	}
	defer db.Close()

	stmt, _:=db.Prepare("insert into usuarios(nom) values(?)")
	stmt.Exec("maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")
)