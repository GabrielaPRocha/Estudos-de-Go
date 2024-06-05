package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:102030@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Corrected query condition: retrieve records with id > 2
	rows, err := db.Query("SELECT id, nome FROM usuarios WHERE id > 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u usuario
		err := rows.Scan(&u.id, &u.nome)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
