package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Connexion *sql.DB
}

func (d *Database) Connex() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic("Connexion à la base de donnée impossible")
	}
	fmt.Println("Connecté ! ")
	d.Connexion = db
}

func (d *Database) Query(request string) {
	d.Connexion.Exec(request)
}
