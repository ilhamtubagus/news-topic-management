package repository

import (
	"github.com/ilhamtubagus/newsTags/domain/entity"
	value "github.com/ilhamtubagus/newsTags/domain/valueObject"
)

type TopicRepository interface {
	SaveTopic(topic *entity.Topic) (*entity.Topic, *value.AppError)
	GetTopic(topic string) (*entity.Topic, *value.AppError)
	GetTopicById(id uint64) (*entity.Topic, *value.AppError)
	DeleteTopic(id uint64) *value.AppError
	GetAllTopic() ([]entity.Topic, *value.AppError)
}
