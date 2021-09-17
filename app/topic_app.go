package app

import (
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/domain/repository"
	value "github.com/ilhamtubagus/newsTags/domain/valueObject"
)

type TopicApp interface {
	SaveTopic(Topic *entity.Topic) (*entity.Topic, *value.AppError)
	GetTopicById(id uint64) (*entity.Topic, *value.AppError)
	GetAllTopic() ([]entity.Topic, *value.AppError)
	DeleteTopic(id uint64) *value.AppError
}
type TopicAppImpl struct {
	TopicRepo repository.TopicRepository
}

func (t *TopicAppImpl) SaveTopic(Topic *entity.Topic) (*entity.Topic, *value.AppError) {
	return t.TopicRepo.SaveTopic(Topic)
}

func (t *TopicAppImpl) GetTopicById(id uint64) (*entity.Topic, *value.AppError) {
	return t.TopicRepo.GetTopicById(id)
}
func (t *TopicAppImpl) DeleteTopic(id uint64) *value.AppError {
	return t.TopicRepo.DeleteTopic(id)
}
func (t *TopicAppImpl) GetAllTopic() ([]entity.Topic, *value.AppError) {
	return t.TopicRepo.GetAllTopic()
}
