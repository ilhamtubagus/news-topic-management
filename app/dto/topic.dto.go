package dto

type TopicDto struct {
	Topic string `json:"topic" validate:"required"`
}
