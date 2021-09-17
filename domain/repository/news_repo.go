package repository

import (
	"github.com/ilhamtubagus/newsTags/domain/entity"
	value "github.com/ilhamtubagus/newsTags/domain/valueObject"
)

type NewsRepository interface {
	SaveNews(*entity.News) (*entity.News, error)
	GetNews(id uint) (*entity.News, error)
	GetAllNews(*value.NewsFilter) (*[]entity.News, error)
	UpdateNews(*entity.News) (*entity.News, error)
	DeleteNews(id uint) error
}
