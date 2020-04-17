package go_sql_hook

import (
	"context"
	"database/sql"
	"testing"
)

func TestHookDriver_Open(t *testing.T) {
	db, _ := sql.Open("hook", "source test")
	tx, _ := db.Begin()
	_, _ = tx.ExecContext(context.Background(), "test query s = ? and b = ?", "1", "2")
}

func TestHookDriver_OpenConnector(t *testing.T) {

}