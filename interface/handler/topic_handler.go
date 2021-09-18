package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ilhamtubagus/newsTags/app"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/infrastructure/persistence"
	"github.com/ilhamtubagus/newsTags/interface/dto"
	"github.com/ilhamtubagus/newsTags/utils"
	"github.com/labstack/echo/v4"
)

type TopicHandler struct {
	topicApp app.TopicApp
	cacher   persistence.Cacher
}

func NewTopicHandler(t app.TopicApp, c persistence.Cacher) *TopicHandler {
	return &TopicHandler{topicApp: t, cacher: c}
}
func (th *TopicHandler) SaveTopic(c echo.Context) error {
	topic := new(entity.Topic)
	if err := c.Bind(topic); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}

	_, errApp := th.topicApp.SaveTopic(topic)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	// flush redis cache
	err := th.cacher.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}
	return c.JSON(http.StatusCreated, dto.TopicDtoRes{Message: "topic created successfully", Topic: topic})
}
func (th *TopicHandler) GetTopicById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "topic not found")

	}
	//check in redis
	key := fmt.Sprintf("topic:%d", id)
	hashedKey := utils.MD5([]byte(key))
	if th.cacher.IsExist(hashedKey) {
		val := th.cacher.Get(hashedKey)
		var topic entity.Topic
		err := json.Unmarshal([]byte(val.(string)), &topic)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "unexpected server error")
		}
		return c.JSON(http.StatusOK, dto.TopicDtoRes{Message: "topic fetched successfully", Topic: &topic})
	}
	// retrieve from database
	topic, errApp := th.topicApp.GetTopicById(id)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())

	}
	// put result in redis
	v, _ := json.Marshal(topic)
	err = th.cacher.Put(hashedKey, v, 600)
	if err != nil {
		c.Logger().Error("unable to put tag in redis cache")
	}
	return c.JSON(http.StatusOK, dto.TopicDtoRes{Message: "topic fetched successfully", Topic: topic})
}
func (th *TopicHandler) UpdateTopic(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "topic not found")
	}
	t := new(entity.Topic)
	if err := c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}
	topic, errApp := th.topicApp.GetTopicById(id)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	topic.Topic = t.Topic
	if _, errApp := th.topicApp.SaveTopic(topic); err != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	// flush redis cache
	err = th.cacher.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}
	return c.JSON(http.StatusOK, dto.TopicDtoRes{Message: "topic updated successfully", Topic: topic})
}
func (th *TopicHandler) DeleteTopic(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "topic not found")
	}
	topic, errApp := th.topicApp.GetTopicById(id)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	errApp = th.topicApp.DeleteTopic(topic.ID)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	// flush redis cache
	err = th.cacher.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}
	return c.JSON(http.StatusOK, dto.TopicDtoRes{Message: "topic deleted successfully"})
}

func (th *TopicHandler) GetAllTopic(c echo.Context) error {
	//check in redis
	hashedKey := utils.MD5([]byte("topics"))
	if th.cacher.IsExist(hashedKey) {
		val := th.cacher.Get(hashedKey)
		var topics []entity.Topic
		err := json.Unmarshal([]byte(val.(string)), &topics)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "unexpected server error")
		}
		return c.JSON(http.StatusOK, dto.TopicsDtoRes{Message: "list of topic fetched successfully", Topics: &topics})
	}
	// retrieve from database
	topics, errApp := th.topicApp.GetAllTopic()
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	v, _ := json.Marshal(topics)
	err := th.cacher.Put(hashedKey, v, 600)
	if err != nil {
		c.Logger().Error("unable to put tag in redis cache")
	}
	return c.JSON(http.StatusOK, dto.TopicsDtoRes{Message: "topic created successfully", Topics: &topics})
}
