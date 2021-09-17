package repository

import (
	"github.com/ilhamtubagus/newsTags/domain/entity"
	value "github.com/ilhamtubagus/newsTags/domain/valueObject"
)

type TagRepository interface {
	SaveTag(*entity.Tag) (*entity.Tag, *value.AppError)
	GetTagById(id uint64) (*entity.Tag, *value.AppError)
	GetTag(tag string) (*entity.Tag, *value.AppError)
	DeleteTag(id uint64) *value.AppError
	GetAllTag() ([]entity.Tag, *value.AppError)
}
