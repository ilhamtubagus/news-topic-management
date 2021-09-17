package handler

import (
	"net/http"
	"strconv"

	"github.com/ilhamtubagus/newsTags/app"
	"github.com/ilhamtubagus/newsTags/domain/entity"
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
	return c.JSON(http.StatusCreated, topic)
}
func (th *TopicHandler) GetTopicById(c echo.Context) error {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		tag, err := th.topicApp.GetTopicById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())

		}
		return c.JSON(http.StatusOK, tag)
	}
	return echo.NewHTTPError(http.StatusNotFound, "tag not found")
}
func (th *TopicHandler) UpdateTopic(c echo.Context) error {
	t := new(entity.Topic)
	if err := c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		tag, err := th.topicApp.GetTopicById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())

		}
		tag.Topic = t.Topic
		if _, err := th.topicApp.SaveTopic(tag); err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())
		}
		return c.JSON(http.StatusOK, tag)
	}
	return echo.NewHTTPError(http.StatusNotFound, "tag not found")
}
func (th *TopicHandler) DeleteTopic(c echo.Context) error {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		tag, err := th.topicApp.GetTopicById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())
		}
		th.topicApp.DeleteTopic(tag.ID)
		return c.JSON(http.StatusOK, "tag deleted")
	}
	return echo.NewHTTPError(http.StatusNotFound, "tag not found")
}

func (th *TopicHandler) GetAllTopic(c echo.Context) error {
	topics, err := th.topicApp.GetAllTopic()
	if err != nil {
		return echo.NewHTTPError(err.Code, err.AsMessage())
	}
	return c.JSON(http.StatusOK, topics)
}
