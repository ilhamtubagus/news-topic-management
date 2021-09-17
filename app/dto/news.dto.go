package dto

type NewsDto struct {
	Title   string
	Author  string
	Status  string
	Content string
	Image   string
	TopicID uint64 `json:"topic_id"`
	Tags    []uint64
}
