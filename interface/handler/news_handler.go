package handler

import (
	"net/http"
	"strconv"

	"github.com/ilhamtubagus/newsTags/app"
	"github.com/ilhamtubagus/newsTags/interface/dto"
	"github.com/labstack/echo/v4"
)

type NewsHandler struct {
	newsApp app.NewsApp
}

func NewNewsandler(n app.NewsApp) *NewsHandler {
	return &NewsHandler{newsApp: n}
}
func (nh *NewsHandler) SaveNews(c echo.Context) error {
	newsDto := new(dto.NewsDto)
	if err := c.Bind(newsDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}
	news, err := nh.newsApp.SaveNews(newsDto)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.AsMessage())
	}
	return c.JSON(http.StatusCreated, dto.NewsDtoRes{Message: "news created", News: news})
}

func (nh *NewsHandler) GetNewsById(c echo.Context) error {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		news, err := nh.newsApp.GetNewsById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())

		}
		return c.JSON(http.StatusOK, dto.NewsDtoRes{Message: "news fetched successfully", News: news})
	}
	return echo.NewHTTPError(http.StatusNotFound, "news not found")
}
func (nh *NewsHandler) GetAllNews(c echo.Context) error {
	var filter dto.NewsFilter
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &filter); err != nil {
		return err
	}
	listOfNews, err := nh.newsApp.GetAllNews(&filter)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.AsMessage())
	}
	return c.JSON(http.StatusOK, dto.ListOfNewsDtoRes{Message: "list of news fetched successfully", News: listOfNews})
}

func (nh *NewsHandler) UpdateNews(c echo.Context) error {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		news, err := nh.newsApp.GetNewsById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())

		}
		newsDto := new(dto.NewsDto)
		if err := c.Bind(newsDto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
		}
		newsDto.ID = news.ID
		// update news
		news, err = nh.newsApp.SaveNews(newsDto)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())
		}
		return c.JSON(http.StatusOK, dto.NewsDtoRes{Message: "news updated successfully", News: news})
	}
	return echo.NewHTTPError(http.StatusNotFound, "news not found")
}
func (nh *NewsHandler) DeleteNews(c echo.Context) error {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		_, err := nh.newsApp.DeleteNews(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())

		}
		return c.JSON(http.StatusOK, dto.NewsDtoRes{Message: "news deleted successfully"})
	}
	return echo.NewHTTPError(http.StatusNotFound, "news not found")
}
