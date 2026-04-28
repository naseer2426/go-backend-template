package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init establishes a GORM connection to Postgres using the given DSN.
func Init(dsn string) error {
	if dsn == "" {
		return fmt.Errorf("database DSN is empty")
	}
	if DB != nil {
		return nil
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = database
	return nil
}

// GetDB returns the global DB handle. Panics if [Init] was not called successfully first.
func GetDB() *gorm.DB {
	if DB == nil {
		panic("db not initialized; call Init with a non-empty DSN first")
	}
	return DB
}
