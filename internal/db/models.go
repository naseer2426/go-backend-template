package db

import (
	"time"

	"gorm.io/gorm"
)

// ExampleItem illustrates a minimal GORM model that matches migrations/000001_create_example_items.*.
//
// Replace ExampleItem (rename the type and table via TableName()), or add more models beside it.
// For every schema change: add a new migration pair (.up/.down), then update structs accordingly.
type ExampleItem struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement"`
	Title     string         `gorm:"column:title;size:255;not null"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

// TableName keeps the plural table explicit; omit if you rely on GORM’s default pluralization.
func (ExampleItem) TableName() string {
	return "example_items"
}
