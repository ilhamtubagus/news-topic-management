package persistence

import (
	"fmt"

	"github.com/ilhamtubagus/newsTags/domain/entity"
	value "github.com/ilhamtubagus/newsTags/domain/valueObject"
	"gorm.io/gorm"
)

type TopicRepoImpl struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) *TopicRepoImpl {
	return &TopicRepoImpl{db: db}
}
func (r *TopicRepoImpl) SaveTopic(topic *entity.Topic) (*entity.Topic, *value.AppError) {
	err := r.db.Save(topic).Error
	if err != nil {
		return nil, value.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return topic, nil
}
func (r *TopicRepoImpl) GetTopic(t string) (*entity.Topic, *value.AppError) {
	var topic entity.Topic
	err := r.db.Where("topic = ?", t).First(&topic).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, value.NewNotFoundError("tag not found")
		}
		return nil, value.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return &topic, nil
}
func (r *TopicRepoImpl) GetTopicById(id uint64) (*entity.Topic, *value.AppError) {
	var topic entity.Topic
	err := r.db.First(&topic, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, value.NewNotFoundError("tag not found")
		}
		return nil, value.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return &topic, nil
}
func (r *TopicRepoImpl) DeleteTopic(id uint64) *value.AppError {
	if err := r.db.Delete(&entity.Topic{}, id).Error; err != nil {
		return value.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return nil
}
func (r *TopicRepoImpl) GetAllTopic() ([]entity.Topic, *value.AppError) {
	var topics []entity.Topic
	err := r.db.Find(&topics).Error
	if err != nil {
		return nil, value.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return topics, nil
}
