package persistence

import (
	"github.com/ilhamtubagus/newsTags/app/dto"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"gorm.io/gorm"
)

type NewsRepoImpl struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepoImpl {
	return &NewsRepoImpl{db: db}
}

func (n NewsRepoImpl) SaveNews(news *entity.News) (*entity.News, *dto.AppError) {
	return nil, nil
}
