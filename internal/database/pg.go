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
func Migrate() {
	_, err := DB.Exec(`ALTER TABLE users ADD COLUMN phone VARCHAR(50)`)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully added password column to users table")
	}
	_, err = DB.Exec(`ALTER TABLE users ADD COLUMN address VARCHAR(200)`)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully added address column to users table")
	}
	_, err = DB.Exec(`
    CREATE TABLE items (
      id SERIAL PRIMARY KEY,
      amount INTEGER NOT NULL,
      happened_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
      create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    )
  `)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully created items table")
	}
	// _, err = DB.Exec(`ALTER TABLE items ALTER COLUMN happened_at TYPE TIMESTAMP`)
	// handleErr(err)
	// log.Println("Successfully changed happened_at column type to TIMESTAMP")
}
func Crud() {
	// CREATE
	_, err := DB.Exec(`INSERT INTO users (email) VALUES ('1@qq.com')`)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully created user")
	}
	// UPDATE
	_, err = DB.Exec(`Update users SET phone = '13556551111' WHERE id = 1`)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully updated user")
	}
	// READ
	// result, err := DB.Query(`SELECT phone FROM users WHERE email = '1@qq.com'`)
	// result, err := DB.Query(`SELECT phone FROM users WHERE email = '1@qq.com' offset 0 limit 3`)
	stmt, err := DB.Prepare("SELECT phone FROM users WHERE email = $1 offset $2 limit $3")
	if err != nil {
		log.Println(err)
	}
	result, err := stmt.Query("1@qq.com", 0, 3)
	if err != nil {
		log.Println(err)
	} else {
		for result.Next() {
			var phone string
			result.Scan(&phone)
			log.Println("phone", phone)
		}
		log.Println("Successfully read a users")
	}
}
func Close() {
	DB.Close()
	log.Println("Successfully closed database connection")
}
