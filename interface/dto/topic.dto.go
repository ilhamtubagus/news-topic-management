package dto

import "github.com/ilhamtubagus/newsTags/domain/entity"

type TopicDto struct {
	Topic string `json:"topic" validate:"required"`
}
type TopicDtoRes struct {
	Message string        `json:"message"`
	Topic   *entity.Topic `json:"topic,omitempty"`
}
type TopicsDtoRes struct {
	Message string          `json:"message"`
	Topics  *[]entity.Topic `json:"topics"`
}
