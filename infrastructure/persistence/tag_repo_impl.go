package persistence

import (
	"fmt"

	"github.com/ilhamtubagus/newsTags/domain/entity"
	value "github.com/ilhamtubagus/newsTags/domain/valueObject"
	"gorm.io/gorm"
)

type TagRepoImpl struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepoImpl {
	return &TagRepoImpl{db: db}
}
func (r *TagRepoImpl) SaveTag(tag *entity.Tag) (*entity.Tag, *value.AppError) {
	err := r.db.Save(tag).Error
	if err != nil {
		return nil, value.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return tag, nil
}
func (r *TagRepoImpl) GetTag(t string) (*entity.Tag, *value.AppError) {
	var tag entity.Tag
	err := r.db.Where("tag = ?", t).First(&tag).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, value.NewNotFoundError("tag not found")
		}
		return nil, value.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return &tag, nil
}

func (r *TagRepoImpl) GetTagById(id uint64) (*entity.Tag, *value.AppError) {
	var tag entity.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, value.NewNotFoundError("tag not found")
		}
		return nil, value.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return &tag, nil
}
func (r *TagRepoImpl) DeleteTag(id uint64) *value.AppError {
	if err := r.db.Delete(&entity.Tag{}, id).Error; err != nil {
		return value.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return nil
}
func (r *TagRepoImpl) GetAllTag() ([]entity.Tag, *value.AppError) {
	var tags []entity.Tag
	err := r.db.Find(&tags).Error
	if err != nil {
		return nil, value.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return tags, nil
}
