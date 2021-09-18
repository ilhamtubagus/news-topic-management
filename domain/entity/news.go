package entity

import (
	"time"
)

type News struct {
	ID          uint64     `gorm:"primaryKey;auto_increment" json:"id"`
	Title       string     `gorm:"not null;" json:"title"`
	Author      string     `gorm:"not null;" json:"author"`
	Status      string     `gorm:"not null;default:draft" json:"status"`
	Content     string     `gorm:"not null;" json:"content"`
	TopicID     uint64     `json:"-"`
	Topic       *Topic     `gorm:"not null;foreignKey:TopicID;" json:"topic"`
	Tags        []Tag      `gorm:"many2many:news_tag;" json:"tags"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	DeletedAt   *time.Time `gorm:"default:NULL" json:"deleted_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	PublishedAt *time.Time `gorm:"default:NULL" json:"published_at"`
}
