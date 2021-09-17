package persistence

import (
	"fmt"

	"github.com/ilhamtubagus/newsTags/app/dto"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"gorm.io/gorm"
)

type TopicRepoImpl struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) *TopicRepoImpl {
	return &TopicRepoImpl{db: db}
}
func (r *TopicRepoImpl) SaveTopic(topic *entity.Topic) (*entity.Topic, *dto.AppError) {
	err := r.db.Save(topic).Error
	if err != nil {
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return topic, nil
}
func (r *TopicRepoImpl) GetTopic(t string) (*entity.Topic, *dto.AppError) {
	var topic entity.Topic
	err := r.db.Where("topic = ?", t).First(&topic).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, dto.NewNotFoundError("topic not found")
		}
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return &topic, nil
}
func (r *TopicRepoImpl) GetTopicById(id uint64) (*entity.Topic, *dto.AppError) {
	var topic entity.Topic
	err := r.db.First(&topic, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, dto.NewNotFoundError("topic not found")
		}
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return &topic, nil
}
func (r *TopicRepoImpl) DeleteTopic(id uint64) *dto.AppError {
	if err := r.db.Delete(&entity.Topic{}, id).Error; err != nil {
		return dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return nil
}
func (r *TopicRepoImpl) GetAllTopic() ([]entity.Topic, *dto.AppError) {
	var topics []entity.Topic
	err := r.db.Find(&topics).Error
	if err != nil {
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return topics, nil
}
