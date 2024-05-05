package database

import (
	"database/sql"
	"fmt"
	"httpserveur/pkg/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Connexion *sql.DB
	Host      string
	Port      string
	User      string
	Password  string
	Dbname    string
}

func NewDatabase() *Database {
	return &Database{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "",
		Dbname:   "golang",
	}
}

func (d *Database) Connex() error {
	db, err := sql.Open("mysql", config.DbConnect)
	if err != nil {
		log.Fatalf("Échec de la connexion à la base de données: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Échec lors de la tentative de connexion à la base de données: %v", err)
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
