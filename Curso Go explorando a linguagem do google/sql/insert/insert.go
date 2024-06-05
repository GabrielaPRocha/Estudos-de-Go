package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:102030@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) value(?)")
	stmt.Exec("Sara Lance")
	stmt.Exec("Barry Quenn")
	stmt.Exec("Lena Luthor")
	stmt.Exec("Tea Snow")
	stmt.Exec("James ")
	stmt.Exec("Mart")

	res, _ := stmt.Exec("Pedro")
	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
