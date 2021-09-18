package entity

import "time"

type Tag struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Tag       string    `gorm:"uniqueIndex;not null" json:"tag,omitempty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}
