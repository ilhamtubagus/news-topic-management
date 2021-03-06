package dto

import "github.com/ilhamtubagus/newsTags/domain/entity"

type TagDto struct {
	Tag string `json:"tag" validate:"required"`
}
type TagDtoRes struct {
	Message string      `json:"message"`
	Tag     *entity.Tag `json:"tag,omitempty"`
}
type TagsDtoRes struct {
	Message string        `json:"message"`
	Tags    *[]entity.Tag `json:"tags"`
}
