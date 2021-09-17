package app

import (
	"github.com/ilhamtubagus/newsTags/app/dto"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/domain/repository"
)

type NewsApp interface {
	SaveNews(n *dto.NewsDto) (*entity.News, *dto.AppError)
}
type NewsAppImpl struct {
	NewsRepo  repository.NewsRepository
	TagRepo   repository.TagRepository
	TopicRepo repository.TopicRepository
}

func (n NewsAppImpl) SaveNews(newsDto *dto.NewsDto) (*entity.News, *dto.AppError) {
	var tags []entity.Tag
	for _, v := range newsDto.Tags {
		tag, err := n.TagRepo.GetTagById(v)
		if err != nil {
			return nil, err
		}
		tags = append(tags, *tag)
	}
	topic, err := n.TopicRepo.GetTopicById(newsDto.TopicID)
	if err != nil {
		return nil, err

	}
	// instantiate news
	news := entity.News{
		Title:   newsDto.Title,
		Author:  newsDto.Author,
		Status:  newsDto.Status,
		Content: newsDto.Content,
		Image:   newsDto.Image,
		Topic:   topic,
		Tags:    tags,
	}
	_, err = n.NewsRepo.SaveNews(&news)
	if err != nil {
		return nil, err
	}
	return &news, nil

}
