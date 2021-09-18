package repository

import (
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/interface/dto"
)

type TagRepository interface {
	SaveTag(*entity.Tag) (*entity.Tag, *dto.AppError)
	GetTagById(id uint64) (*entity.Tag, *dto.AppError)
	GetTag(tag string) (*entity.Tag, *dto.AppError)
	DeleteTag(id uint64) *dto.AppError
	GetAllTag() ([]entity.Tag, *dto.AppError)
}
