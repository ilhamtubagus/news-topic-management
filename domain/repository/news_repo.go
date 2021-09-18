package repository

import (
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/interface/dto"
)

type NewsRepository interface {
	SaveNews(*entity.News) (*entity.News, *dto.AppError)
	GetNewsById(id uint64) (*entity.News, *dto.AppError)
	GetAllNews(*dto.NewsFilter) (*[]entity.News, *dto.AppError)
	// UpdateNews(*entity.News) (*entity.News, error)
	// DeleteNews(id uint) error
}
