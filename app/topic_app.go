package app

import (
	"github.com/ilhamtubagus/newsTags/app/dto"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/domain/repository"
)

type TopicApp interface {
	SaveTopic(Topic *entity.Topic) (*entity.Topic, *dto.AppError)
	GetTopicById(id uint64) (*entity.Topic, *dto.AppError)
	GetAllTopic() ([]entity.Topic, *dto.AppError)
	DeleteTopic(id uint64) *dto.AppError
}
type TopicAppImpl struct {
	TopicRepo repository.TopicRepository
}

func (t *TopicAppImpl) SaveTopic(Topic *entity.Topic) (*entity.Topic, *dto.AppError) {
	return t.TopicRepo.SaveTopic(Topic)
}

func (t *TopicAppImpl) GetTopicById(id uint64) (*entity.Topic, *dto.AppError) {
	return t.TopicRepo.GetTopicById(id)
}
func (t *TopicAppImpl) DeleteTopic(id uint64) *dto.AppError {
	return t.TopicRepo.DeleteTopic(id)
}
func (t *TopicAppImpl) GetAllTopic() ([]entity.Topic, *dto.AppError) {
	return t.TopicRepo.GetAllTopic()
}
