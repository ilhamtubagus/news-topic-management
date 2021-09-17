package persistence

import (
	"fmt"

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
	err := n.db.Save(news).Error
	if err != nil {
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return news, nil
}
