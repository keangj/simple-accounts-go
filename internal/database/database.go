package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"simple-accounts/config/tutorial"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB
var DBCtx = context.Background()

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "123456"
	dbname   = "simple_accounts_dev"
)

type User struct {
	ID        int
	Email     string `gorm:"uniqueIndex"` // 使用 tag 设置唯一索引
	Phone     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Items struct {
	ID        int
	TagId     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Tag struct {
	ID int
}

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	err = DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

var models = []any{&User{}, &Items{}, &Tag{}}

func CreateTables() {

}
func CreateMigrate(filename string) {
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "config/migrations", "-seq", filename)
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
func Migrate() {
	dir, err := os.Getwd()
	name := filepath.Base(dir)
	for !strings.Contains(name, "simple-accounts-go") {
		dir = filepath.Dir(dir)
		name = filepath.Base(dir)
	}
	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", dir),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname),
	)
	if err != nil {
		log.Fatalln(err)
	}
	// m.Steps(1) // 迁移一步
	err = m.Up() // 迁移所有
	if err != nil {
		if !strings.Contains(err.Error(), "no change") {
			log.Fatalln(err)
		}
	}
}

func MigrateDown() {
	dir, err := os.Getwd()
	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", dir),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname),
	)
	if err != nil {
		log.Fatalln(err)
	}
	// err = m.Down() // 回滚所有
	err = m.Steps(-1) // 回滚一步
	if err != nil {
		log.Fatalln(err)
	}
}

func Crud() {
	// Create
	t := tutorial.New(DB)
	id := rand.Int()
	u, err := t.CreateUser(DBCtx, fmt.Sprintf("%d@qq.com", id))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(u)
	// Update
	t.UpdateUser(DBCtx, tutorial.UpdateUserParams{
		ID:      u.ID,
		Email:   u.Email,
		Phone:   u.Phone,
		Address: "北京",
	})
	// Read
	users, err := t.ListUsers(DBCtx, tutorial.ListUsersParams{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(users)
	u, err = t.GetUser(DBCtx, users[0].ID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(u)
	// Delete
	err = t.DeleteUser(DBCtx, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
}
func Close() {

}

// ------ 以下为 GORM 写法 ------
// package database

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "admin"
// 	password = "123456"
// 	dbname   = "simple_accounts_dev"
// )

// type User struct {
// 	ID        int
// 	Email     string `gorm:"uniqueIndex"` // 使用 tag 设置唯一索引
// 	Phone     int
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
// type Items struct {
// 	ID        int
// 	TagId     string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
// type Tag struct {
// 	ID int
// }

// func Connect() {
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	// 连接数据库
// 	// gorm 默认自动 Ping 不需自动 Ping 需要在 &gorm.Config{} 中设置 DisableAutomaticPing 为 true
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	DB = db
// }

// var models = []any{&User{}, &Items{}, &Tag{}}

// func CreateTables() {
// 	for _, model := range models {
// 		err := DB.Migrator().CreateTable(model)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 	}
// 	log.Println("Successfully created table")
// }
// func Migrate() {
// 	err := DB.AutoMigrate(models...)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	log.Println("Successfully migrated table")
// }
// func Crud() {
// 	// Create
// 	user := User{Email: "5@qq.com"}
// 	tx := DB.Create(&user)
// 	log.Println(tx.RowsAffected) // tx 是一个事务，可以获取事务的执行结果
// 	log.Println(user)            // user 的值会变为新创建的值

// 	// Read
// 	// 获取多个
// 	users := []User{}
// 	// 查询所有的 users
// 	// DB.Find(&users)
// 	// 查询主键为 1,2,3 的 users
// 	DB.Find(&users, []int{1, 2, 3})
// 	// 查询前 10 个 users
// 	// asc 升序 desc 降序
// 	DB.Offset(0).Limit(10).Order("created_at asc, id desc").Find(&users)
// 	log.Println(users)
// 	// 获取一个
// 	user = User{}
// 	DB.Find(&user, 1) // 查询主键为 1 的 user
// 	log.Println(user) // user 的值会变为查询的结果

// 	// Update
// 	user.Phone = 123456789
// 	DB.Save(&user)
// 	log.Println(user)

// 	// Delete
// 	// DB.Delete(&User{ID: 1})
// }
// func Close() {
// 	sqlDB, err := DB.DB()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	sqlDB.Close()
// }

// ----- 以下为原生 sql 的写法 -----
// package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	"github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "admin"
// 	password = "123456"
// 	dbname   = "simple_accounts_dev"
// )

// func Connect() {
// 	connectStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", connectStr)
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

// func CreateTables() {
// 	_, err := DB.Exec(`
// 		CREATE TABLE IF NOT EXISTS users (
// 			id SERIAL PRIMARY KEY,
// 			email VARCHAR(100) NOT NULL,
// 			create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 			update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
// 		)`)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }
// func Migrate() {
// 	_, err := DB.Exec(`ALTER TABLE users ADD COLUMN phone VARCHAR(50)`)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	_, err = DB.Exec(`ALTER TABLE users ADD COLUMN address VARCHAR(200)`)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	_, err = DB.Exec(`
//     CREATE TABLE items (
//       id SERIAL PRIMARY KEY,
//       amount INTEGER NOT NULL,
//       happened_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
//       create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//       update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
//     )
//   `)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	_, err = DB.Exec(`ALTER TABLE items ALTER COLUMN happened_at TYPE TIMESTAMP`)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	// 为 users 的 email 字段添加唯一索引
// 	_, err = DB.Exec(`CREATE UNIQUE INDEX users_email_index ON users (email)`)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
// func Crud() {
// 	// CREATE
// 	_, err := DB.Exec(`INSERT INTO users (email) VALUES ('1@qq.com')`)
// 	if err != nil {
// 		switch x := err.(type) {
// 		case *pq.Error:
// 			pqErr := err.(*pq.Error)
// 			log.Println(pqErr.Code.Name())
// 			log.Panicln(pqErr.Message)
// 		default:
// 			log.Println(x)
// 		}
// 	}
// 	// UPDATE
// 	_, err = DB.Exec(`Update users SET phone = '13556551111' WHERE id = 1`)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	// READ
// 	// result, err := DB.Query(`SELECT phone FROM users WHERE email = '1@qq.com'`)
// 	// result, err := DB.Query(`SELECT phone FROM users WHERE email = '1@qq.com' offset 0 limit 3`)
// 	stmt, err := DB.Prepare("SELECT phone FROM users WHERE email = $1 offset $2 limit $3")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	result, err := stmt.Query("1@qq.com", 0, 3)
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		for result.Next() {
// 			var phone string
// 			result.Scan(&phone)
// 			log.Println("phone", phone)
// 		}
// 	}
// }
// func Close() {
// 	DB.Close()
// 	log.Println("Successfully closed database connection")
// }
