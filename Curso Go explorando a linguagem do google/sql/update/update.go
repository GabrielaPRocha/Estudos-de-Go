package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:102030@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//update
	/*stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Karl Zoerl", 1)
	stmt.Exec("Oliver Allen", 2)]*/

	//delete
	stmt, _ := db.Prepare("delete from usuarios where id = ?")
	stmt.Exec(3)

}
