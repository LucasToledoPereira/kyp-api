package entities

import "time"

type CustomMigrations struct {
	ID        int `gorm:"primaryKey; index; unique;"`
	Name      string
	CreatedAt time.Time
}
