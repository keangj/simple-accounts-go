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
func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
func Migrate() {
	_, err := DB.Exec(`ALTER TABLE users ADD COLUMN phone VARCHAR(50)`)
	handleErr(err)
	log.Println("Successfully added password column to users table")
	_, err = DB.Exec(`ALTER TABLE users ADD COLUMN address VARCHAR(200)`)
	handleErr(err)
	log.Println("Successfully added address column to users table")
	_, err = DB.Exec(`
    CREATE TABLE items (
      id SERIAL PRIMARY KEY,
      amount INTEGER NOT NULL,
      happened_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
      create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    )
  `)
	handleErr(err)
	log.Println("Successfully created items table")
	// _, err = DB.Exec(`ALTER TABLE items ALTER COLUMN happened_at TYPE TIMESTAMP`)
	// handleErr(err)
	// log.Println("Successfully changed happened_at column type to TIMESTAMP")
}
func Close() {
	DB.Close()
	log.Println("Successfully closed database connection")
}
