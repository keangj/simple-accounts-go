package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// const (
// 	mysqlHost     = "localhost"
// 	mysqlPort     = 3306
// 	mysqlUser     = "jay"
// 	mysqlPassword = "123456"
// 	mysqlDbname   = "simple_accounts_dev"
// )

// func MysqlConnect() {
// 	connectStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDbname)
// 	db, err := sql.Open("mysql", connectStr)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	DB = db
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	log.Println("Successfully connected to database")
// }

// func MysqlCreateTables() {
// 	_, err := DB.Exec(`
// 		CREATE TABLE IF NOT EXISTS users (
// 			id INT AUTO_INCREMENT PRIMARY KEY,
// 			email VARCHAR(50) NOT NULL,
// 			create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 			update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 		)`)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	log.Println("Successfully created users table")
// }

// func MysqlClose() {
// 	DB.Close()
// 	log.Println("Successfully closed database connection")
// }
