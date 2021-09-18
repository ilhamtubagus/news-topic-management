package dto

import "github.com/ilhamtubagus/newsTags/domain/entity"

type NewsDto struct {
	ID      uint64   `json:"id,omitempty"`
	Title   string   `json:"title,omitempty"`
	Author  string   `json:"author,omitempty"`
	Status  string   `json:"status,omitempty"`
	Content string   `json:"content,omitempty"`
	TopicID uint64   `json:"topic_id,omitempty"`
	Tags    []uint64 `json:"tags,omitempty"`
}
type NewsDtoRes struct {
	Message string       `json:"message"`
	News    *entity.News `json:"news,omitempty"`
}
type ListOfNewsDtoRes struct {
	Message string         `json:"message"`
	News    *[]entity.News `json:"news"`
}
