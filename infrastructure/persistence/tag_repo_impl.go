package persistence

import (
	"fmt"

	"github.com/ilhamtubagus/newsTags/app/dto"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"gorm.io/gorm"
)

type TagRepoImpl struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepoImpl {
	return &TagRepoImpl{db: db}
}
func (r *TagRepoImpl) SaveTag(tag *entity.Tag) (*entity.Tag, *dto.AppError) {
	err := r.db.Save(tag).Error
	if err != nil {
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return tag, nil
}
func (r *TagRepoImpl) GetTag(t string) (*entity.Tag, *dto.AppError) {
	var tag entity.Tag
	err := r.db.Where("tag = ?", t).First(&tag).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, dto.NewNotFoundError("tag not found")
		}
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return &tag, nil
}

func (r *TagRepoImpl) GetTagById(id uint64) (*entity.Tag, *dto.AppError) {
	var tag entity.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, dto.NewNotFoundError("tag not found")
		}
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return &tag, nil
}
func (r *TagRepoImpl) DeleteTag(id uint64) *dto.AppError {
	if err := r.db.Delete(&entity.Tag{}, id).Error; err != nil {
		return dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return nil
}
func (r *TagRepoImpl) GetAllTag() ([]entity.Tag, *dto.AppError) {
	var tags []entity.Tag
	err := r.db.Find(&tags).Error
	if err != nil {
		return nil, dto.NewUnexpectedError(fmt.Sprintf("unexpected database error [%s]", err.Error()))
	}
	return tags, nil
}
