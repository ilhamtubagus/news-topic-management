package entity

import (
	"time"

	value "github.com/ilhamtubagus/newsTags/domain/valueObject"
)

type News struct {
	ID          uint64 `gorm:"primaryKey;auto_increment"`
	Title       string `gorm:"notnull;"`
	Author      string `gorm:"notnull;"`
	Status      string `gorm:"notnull;"`
	Content     string `gorm:"notnull;"`
	Image       string
	TopicID     uint64
	Topic       *Topic     `gorm:"notnull;foreignKey:TopicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tags        []Tag      `gorm:"many2many:news_tag;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	DeletedAt   *time.Time `gorm:"default:NULL" json:"deleted_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	PublishedAt *time.Time `gorm:"default:NULL" json:"published_at"`
}

func (n News) Validate() bool {
	return n.Status == value.DeletedNews || n.Status == value.DraftedNews || n.Status == value.PublishedNews
}
