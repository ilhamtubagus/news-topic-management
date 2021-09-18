package app

import (
	"time"

	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/domain/repository"
	"github.com/ilhamtubagus/newsTags/interface/dto"
)

type NewsApp interface {
	SaveNews(n *dto.NewsDto) (*entity.News, *dto.AppError)
	GetNewsById(id uint64) (*entity.News, *dto.AppError)
	GetAllNews(filter *dto.NewsFilter) (*[]entity.News, *dto.AppError)
	DeleteNews(id uint64) (*entity.News, *dto.AppError)
}
type NewsAppImpl struct {
	NewsRepo  repository.NewsRepository
	TagRepo   repository.TagRepository
	TopicRepo repository.TopicRepository
}

func (n NewsAppImpl) SaveNews(newsDto *dto.NewsDto) (*entity.News, *dto.AppError) {
	//domain validation
	if newsDto.Status != "draft" && newsDto.Status != "published" && newsDto.Status != "deleted" {
		return nil, dto.NewBadRequestError("status not valid")
	}
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
		ID:      newsDto.ID,
		Title:   newsDto.Title,
		Author:  newsDto.Author,
		Status:  newsDto.Status,
		Content: newsDto.Content,
		Topic:   topic,
		Tags:    tags,
	}
	if news.Status == "published" {
		now := time.Now()
		news.PublishedAt = &now
	}
	_, err = n.NewsRepo.SaveNews(&news)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (n NewsAppImpl) GetNewsById(id uint64) (*entity.News, *dto.AppError) {
	return n.NewsRepo.GetNewsById(id)
}
func (n NewsAppImpl) GetAllNews(filter *dto.NewsFilter) (*[]entity.News, *dto.AppError) {
	return n.NewsRepo.GetAllNews(filter)
}
func (n NewsAppImpl) DeleteNews(id uint64) (*entity.News, *dto.AppError) {
	news, err := n.NewsRepo.GetNewsById(id)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	news.DeletedAt = &now
	news.Status = "deleted"
	news, err = n.NewsRepo.SaveNews(news)
	if err != nil {
		return nil, err
	}
	return news, nil
}
