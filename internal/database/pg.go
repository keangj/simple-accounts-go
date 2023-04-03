package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "123456"
	dbname   = "simple_accounts_dev"
)

func Connect() {
	connectStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully connected to database")
}

func CreateTables() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(100) NOT NULL,
			create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully created users table")
}

func Close() {
	DB.Close()
	log.Println("Successfully closed database connection")
}
