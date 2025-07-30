package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sqlx.DB

func Connect() error {
	dsn := "root:root@tcp(127.0.0.1:3307)/sykell?parseTime=true"
	var err error
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	return nil
}
