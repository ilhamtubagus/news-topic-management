package entity

import "time"

type Tag struct {
	ID        uint64    `gorm:"primaryKey;auto_increment" json:"id,omitempty"`
	Tag       string    `gorm:"uniqueIndex" json:"tag,omitempty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func (t Tag) Validate() bool {
	return t.Tag == ""
}
