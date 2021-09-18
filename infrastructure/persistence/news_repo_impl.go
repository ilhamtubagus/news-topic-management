package persistence

import (
	"fmt"

	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/interface/dto"
	"gorm.io/gorm"
)

type NewsRepoImpl struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepoImpl {
	return &NewsRepoImpl{db: db}
}

func (n NewsRepoImpl) SaveNews(news *entity.News) (*entity.News, *dto.AppError) {
	err := n.db.Save(news).Error
	if err != nil {
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return news, nil
}

func (n NewsRepoImpl) GetNewsById(id uint64) (*entity.News, *dto.AppError) {
	var news entity.News
	err := n.db.Where("id = ?", id).Preload("Topic").Preload("Tags").Find(&news).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, dto.NewNotFoundError("news not found")
		}
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return &news, nil
}

func (n NewsRepoImpl) GetAllNews(filter *dto.NewsFilter) (*[]entity.News, *dto.AppError) {
	exec := n.db.Preload("Topic").Preload("Tags")
	var news []entity.News

	var err error
	if filter.Status != "" {
		if filter.Topic != 0 {
			err = exec.Find(&news, "status = ? and topic_id = ?", filter.Status, filter.Topic).Error
		} else {
			err = exec.Find(&news, "status = ?", filter.Status).Error
		}
	} else {
		if filter.Topic != 0 {
			err = exec.Find(&news, "topic_id = ?", filter.Topic).Error
		} else {
			err = exec.Find(&news).Error
		}
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, dto.NewNotFoundError("news not found")
		}
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return &news, nil

}
