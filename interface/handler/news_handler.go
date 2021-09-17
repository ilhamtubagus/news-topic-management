package handler

import (
	"net/http"

	"github.com/ilhamtubagus/newsTags/app"
	"github.com/ilhamtubagus/newsTags/app/dto"
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
	return c.JSON(http.StatusCreated, news)
}
