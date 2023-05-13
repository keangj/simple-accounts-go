package database_test

import (
	"simple-accounts/internal/database"
	"testing"
)

func BenchmarkCrud(b *testing.B) {
	database.Connect()
	// database.CreateTables() // sqlc 不需要执行创建表
	database.Migrate()
	defer database.Close()
	for i := 0; i < b.N; i++ {
		database.Crud()
	}
}
