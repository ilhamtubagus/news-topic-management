package entity

import "time"

type Topic struct {
	ID        uint64    `gorm:"primaryKey;auto_increment"`
	Topic     string    `gorm:"uniqueIndex"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
