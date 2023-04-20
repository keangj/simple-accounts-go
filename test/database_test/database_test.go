package database_test

import (
	"simple-accounts/internal/database"
	"testing"
)

func BenchmarkCrud(b *testing.B) {
	database.Connect()
	database.CreateTables()
	database.Migrate()
	for i := 0; i < b.N; i++ {
		database.Crud()
	}
}
