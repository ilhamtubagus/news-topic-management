package handler

import (
	"net/http"
	"strconv"

	"github.com/ilhamtubagus/newsTags/app"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/interface/dto"
	"github.com/labstack/echo/v4"
)

type TopicHandler struct {
	topicApp app.TopicApp
}

func NewTopicHandler(t app.TopicApp) *TopicHandler {
	return &TopicHandler{topicApp: t}
}
func (th *TopicHandler) SaveTopic(c echo.Context) error {
	topic := new(entity.Topic)
	if err := c.Bind(topic); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}

	_, err := th.topicApp.SaveTopic(topic)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.AsMessage())
	}
	return c.JSON(http.StatusCreated, dto.TopicDtoRes{Message: "topic created successfully", Topic: topic})
}
func (th *TopicHandler) GetTopicById(c echo.Context) error {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		topic, err := th.topicApp.GetTopicById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())

		}
		return c.JSON(http.StatusOK, dto.TopicDtoRes{Message: "topic fetched successfully", Topic: topic})
	}
	return echo.NewHTTPError(http.StatusNotFound, "topic not found")
}
func (th *TopicHandler) UpdateTopic(c echo.Context) error {
	t := new(entity.Topic)
	if err := c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		topic, err := th.topicApp.GetTopicById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())

		}
		topic.Topic = t.Topic
		if _, err := th.topicApp.SaveTopic(topic); err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())
		}
		return c.JSON(http.StatusOK, dto.TopicDtoRes{Message: "topic updated successfully", Topic: topic})
	}
	return echo.NewHTTPError(http.StatusNotFound, "topic not found")
}
func (th *TopicHandler) DeleteTopic(c echo.Context) error {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		topic, err := th.topicApp.GetTopicById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())
		}
		err = th.topicApp.DeleteTopic(topic.ID)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())
		}
		return c.JSON(http.StatusOK, dto.TopicDtoRes{Message: "topic deleted successfully"})
	}
	return echo.NewHTTPError(http.StatusNotFound, "topic not found")
}

func (th *TopicHandler) GetAllTopic(c echo.Context) error {
	topics, err := th.topicApp.GetAllTopic()
	if err != nil {
		return echo.NewHTTPError(err.Code, err.AsMessage())
	}
	return c.JSON(http.StatusOK, dto.TopicsDtoRes{Message: "topic created successfully", Topics: &topics})
}
