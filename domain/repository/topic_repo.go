package repository

import (
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/interface/dto"
)

type TopicRepository interface {
	SaveTopic(topic *entity.Topic) (*entity.Topic, *dto.AppError)
	GetTopic(topic string) (*entity.Topic, *dto.AppError)
	GetTopicById(id uint64) (*entity.Topic, *dto.AppError)
	DeleteTopic(id uint64) *dto.AppError
	GetAllTopic() ([]entity.Topic, *dto.AppError)
}
