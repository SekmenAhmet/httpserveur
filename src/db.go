package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Connexion *sql.DB
}

func (d *Database) Connex() error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		return fmt.Errorf("échec de la connexion à la base de données: %v", err)
	}
	if err = db.Ping(); err != nil {
		return fmt.Errorf("échec lors de la tentative de connexion à la base de données: %v", err)
	}
	fmt.Println("Connecté à la base de données !")
	d.Connexion = db
	return nil
}

func (d *Database) Query(request string, params ...interface{}) error {
	if d.Connexion == nil {
		return fmt.Errorf("connexion à la base de données non initialisée")
	}

	_, err := d.Connexion.Exec(request, params...)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête : %v", err)
		return err
	}
	return nil
}
