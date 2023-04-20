package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
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
}
func Migrate() {
	_, err := DB.Exec(`ALTER TABLE users ADD COLUMN phone VARCHAR(50)`)
	if err != nil {
		log.Println(err)
	}
	_, err = DB.Exec(`ALTER TABLE users ADD COLUMN address VARCHAR(200)`)
	if err != nil {
		log.Println(err)
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
	}
	_, err = DB.Exec(`ALTER TABLE items ALTER COLUMN happened_at TYPE TIMESTAMP`)
	if err != nil {
		log.Println(err)
	}
	// 为 users 的 email 字段添加唯一索引
	_, err = DB.Exec(`CREATE UNIQUE INDEX users_email_index ON users (email)`)
	if err != nil {
		log.Println(err)
	}
}
func Crud() {
	// CREATE
	_, err := DB.Exec(`INSERT INTO users (email) VALUES ('1@qq.com')`)
	if err != nil {
		switch x := err.(type) {
		case *pq.Error:
			pqErr := err.(*pq.Error)
			log.Println(pqErr.Code.Name())
			log.Panicln(pqErr.Message)
		default:
			log.Println(x)
		}
	}
	// UPDATE
	_, err = DB.Exec(`Update users SET phone = '13556551111' WHERE id = 1`)
	if err != nil {
		log.Println(err)
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
	}
}
func Close() {
	DB.Close()
	log.Println("Successfully closed database connection")
}
