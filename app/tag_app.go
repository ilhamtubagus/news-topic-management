package app

import (
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/domain/repository"
	value "github.com/ilhamtubagus/newsTags/domain/valueObject"
)

type TagApp interface {
	SaveTag(tag *entity.Tag) (*entity.Tag, *value.AppError)
	GetTagById(id uint64) (*entity.Tag, *value.AppError)
	GetAllTag() ([]entity.Tag, *value.AppError)
	DeleteTag(id uint64) *value.AppError
}
type TagAppImpl struct {
	TagRepo repository.TagRepository
}

func (t *TagAppImpl) SaveTag(tag *entity.Tag) (*entity.Tag, *value.AppError) {
	return t.TagRepo.SaveTag(tag)
}

func (t *TagAppImpl) GetTagById(id uint64) (*entity.Tag, *value.AppError) {
	return t.TagRepo.GetTagById(id)
}
func (t *TagAppImpl) DeleteTag(id uint64) *value.AppError {
	return t.TagRepo.DeleteTag(id)
}
func (t *TagAppImpl) GetAllTag() ([]entity.Tag, *value.AppError) {
	return t.TagRepo.GetAllTag()
}
