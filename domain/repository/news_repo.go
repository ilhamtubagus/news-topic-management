package repository

import (
	"github.com/ilhamtubagus/newsTags/app/dto"
	"github.com/ilhamtubagus/newsTags/domain/entity"
)

type NewsRepository interface {
	SaveNews(*entity.News) (*entity.News, *dto.AppError)
	// GetNews(id uint) (*entity.News, error)
	// GetAllNews(*dto.NewsFilter) (*[]entity.News, error)
	// UpdateNews(*entity.News) (*entity.News, error)
	// DeleteNews(id uint) error
}
