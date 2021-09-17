package persistence

import (
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/domain/repository"
	"gorm.io/gorm"
)

type DatabaseRepositories struct {
	db              *gorm.DB
	TagRepository   repository.TagRepository
	TopicRepository repository.TopicRepository
}

func (d *DatabaseRepositories) AutoMigrate() error {
	return d.db.AutoMigrate(&entity.Tag{}, &entity.News{}, &entity.Topic{})
}
func NewDatabaseRepositories(dbClient *gorm.DB) *DatabaseRepositories {
	return &DatabaseRepositories{db: dbClient, TagRepository: NewTagRepository(dbClient), TopicRepository: NewTopicRepository((dbClient))}
}
