package app

import (
	"github.com/ilhamtubagus/newsTags/app/dto"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/domain/repository"
)

type TagApp interface {
	SaveTag(tag *entity.Tag) (*entity.Tag, *dto.AppError)
	GetTagById(id uint64) (*entity.Tag, *dto.AppError)
	GetAllTag() ([]entity.Tag, *dto.AppError)
	DeleteTag(id uint64) *dto.AppError
}
type TagAppImpl struct {
	TagRepo repository.TagRepository
}

func (t *TagAppImpl) SaveTag(tag *entity.Tag) (*entity.Tag, *dto.AppError) {
	return t.TagRepo.SaveTag(tag)
}

func (t *TagAppImpl) GetTagById(id uint64) (*entity.Tag, *dto.AppError) {
	return t.TagRepo.GetTagById(id)
}
func (t *TagAppImpl) DeleteTag(id uint64) *dto.AppError {
	return t.TagRepo.DeleteTag(id)
}
func (t *TagAppImpl) GetAllTag() ([]entity.Tag, *dto.AppError) {
	return t.TagRepo.GetAllTag()
}
