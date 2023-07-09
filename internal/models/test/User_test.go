package models_test

import (
	"testing"

	"github.com/rainbowriverrr/CafeGo/pkg/database"
	_ "modernc.org/sqlite"
)

func TestCreateUserTables(t *testing.T) {
	db := database.InitDB(testDB)

	defer db.Close()
	defer clean()
}
