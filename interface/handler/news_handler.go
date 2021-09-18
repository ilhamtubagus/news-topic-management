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

type NewsHandler struct {
	newsApp app.NewsApp
	cacher  persistence.Cacher
}

func NewNewsHandler(n app.NewsApp, c persistence.Cacher) *NewsHandler {
	return &NewsHandler{newsApp: n, cacher: c}
}
func (nh *NewsHandler) SaveNews(c echo.Context) error {
	newsDto := new(dto.NewsDto)
	if err := c.Bind(newsDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}
	news, errApp := nh.newsApp.SaveNews(newsDto)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	// flush redis cache
	err := nh.cacher.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}
	return c.JSON(http.StatusCreated, dto.NewsDtoRes{Message: "news created", News: news})
}

func (nh *NewsHandler) GetNewsById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "news not found")

	}
	//check in redis
	key := fmt.Sprintf("news:%d", id)
	hashedKey := utils.MD5([]byte(key))
	if nh.cacher.IsExist(hashedKey) {
		val := nh.cacher.Get(hashedKey)
		var news entity.News
		err := json.Unmarshal([]byte(val.(string)), &news)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "unexpected server error")
		}
		return c.JSON(http.StatusOK, dto.NewsDtoRes{Message: "list of news fetched successfully", News: &news})
	}
	// retrieve from database
	news, errApp := nh.newsApp.GetNewsById(id)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())

	}
	// put result in redis
	v, _ := json.Marshal(news)
	err = nh.cacher.Put(hashedKey, v, 600)
	if err != nil {
		c.Logger().Error("unable to put news in redis cache")
	}
	return c.JSON(http.StatusOK, dto.NewsDtoRes{Message: "news fetched successfully", News: news})
}
func (nh *NewsHandler) GetAllNews(c echo.Context) error {
	var filter dto.NewsFilter
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse query params")
	}
	//check in redis
	key := fmt.Sprintf("s:%s,t:%d", filter.Status, filter.Topic)
	hashedKey := utils.MD5([]byte(key))
	if nh.cacher.IsExist(hashedKey) {
		val := nh.cacher.Get(hashedKey)
		var listOfNews []entity.News
		err := json.Unmarshal([]byte(val.(string)), &listOfNews)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "unexpected server error")
		}
		return c.JSON(http.StatusOK, dto.ListOfNewsDtoRes{Message: "list of news fetched successfully", News: &listOfNews})
	}
	// retrieve from database
	listOfNews, errApp := nh.newsApp.GetAllNews(&filter)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	// put result in redis
	v, _ := json.Marshal(listOfNews)
	err := nh.cacher.Put(hashedKey, v, 600)
	if err != nil {
		c.Logger().Error("unable to put list of news in redis cache")
	}
	return c.JSON(http.StatusOK, dto.ListOfNewsDtoRes{Message: "list of news fetched successfully", News: listOfNews})
}

func (nh *NewsHandler) UpdateNews(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "news not found")

	}
	news, errApp := nh.newsApp.GetNewsById(id)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())

	}
	newsDto := new(dto.NewsDto)
	if err := c.Bind(newsDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}
	newsDto.ID = news.ID
	news, errApp = nh.newsApp.SaveNews(newsDto)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	// flush redis cache
	err = nh.cacher.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}
	return c.JSON(http.StatusOK, dto.NewsDtoRes{Message: "news updated successfully", News: news})
}
func (nh *NewsHandler) DeleteNews(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "news not found")

	}
	_, errApp := nh.newsApp.DeleteNews(id)
	if err != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())

	}
	return c.JSON(http.StatusOK, dto.NewsDtoRes{Message: "news deleted successfully"})
}
