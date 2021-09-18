package handler

import (
	"net/http"
	"strconv"

	"github.com/ilhamtubagus/newsTags/app"
	"github.com/ilhamtubagus/newsTags/domain/entity"
	"github.com/ilhamtubagus/newsTags/interface/dto"
	"github.com/labstack/echo/v4"
)

type TagHandler struct {
	tagApp app.TagApp
}

func NewTagHandler(t app.TagApp) *TagHandler {
	return &TagHandler{tagApp: t}
}
func (th *TagHandler) SaveTag(c echo.Context) error {
	tag := new(entity.Tag)
	if err := c.Bind(tag); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}

	_, err := th.tagApp.SaveTag(tag)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.AsMessage())
	}
	return c.JSON(http.StatusCreated, dto.TagDtoRes{Message: "tag created successfully", Tag: tag})
}

func (th *TagHandler) GetTagById(c echo.Context) error {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		tag, err := th.tagApp.GetTagById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())

		}
		return c.JSON(http.StatusOK, dto.TagDtoRes{Message: "tag fetched successfully", Tag: tag})
	}
	return echo.NewHTTPError(http.StatusNotFound, "tag not found")
}

func (th *TagHandler) UpdateTag(c echo.Context) error {
	t := new(entity.Tag)
	if err := c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		tag, err := th.tagApp.GetTagById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())

		}
		tag.Tag = t.Tag
		if _, err := th.tagApp.SaveTag(tag); err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())
		}
		return c.JSON(http.StatusOK, dto.TagDtoRes{Message: "tag updated successfully", Tag: tag})
	}
	return echo.NewHTTPError(http.StatusNotFound, "tag not found")
}

func (th *TagHandler) DeleteTag(c echo.Context) error {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 32); err == nil {
		tag, err := th.tagApp.GetTagById(id)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())
		}
		err = th.tagApp.DeleteTag(tag.ID)
		if err != nil {
			return echo.NewHTTPError(err.Code, err.AsMessage())
		}
		return c.JSON(http.StatusOK, dto.TagDtoRes{Message: "tag deleted successfully"})
	}
	return echo.NewHTTPError(http.StatusNotFound, "tag not found")
}
func (th *TagHandler) GetAllTag(c echo.Context) error {
	tags, err := th.tagApp.GetAllTag()
	if err != nil {
		return echo.NewHTTPError(err.Code, err.AsMessage())
	}
	return c.JSON(http.StatusOK, dto.TagsDtoRes{Message: "tags fetched successfully", Tags: &tags})
}
