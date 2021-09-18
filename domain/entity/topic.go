package entity

import "time"

type Topic struct {
	ID        uint64    `gorm:"primaryKey;auto_increment" json:"id"`
	Topic     string    `gorm:"uniqueIndex;not null" json:"topic"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
