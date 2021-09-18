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

type TagHandler struct {
	tagApp app.TagApp
	cacher persistence.Cacher
}

func NewTagHandler(t app.TagApp, c persistence.Cacher) *TagHandler {
	return &TagHandler{tagApp: t, cacher: c}
}
func (th *TagHandler) SaveTag(c echo.Context) error {
	tag := new(entity.Tag)
	if err := c.Bind(tag); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}

	_, errApp := th.tagApp.SaveTag(tag)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	// flush redis cache
	err := th.cacher.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}
	return c.JSON(http.StatusCreated, dto.TagDtoRes{Message: "tag created successfully", Tag: tag})
}

func (th *TagHandler) GetTagById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "tag not found")
	}
	//check in redis
	key := fmt.Sprintf("tag:%d", id)
	hashedKey := utils.MD5([]byte(key))
	if th.cacher.IsExist(hashedKey) {
		val := th.cacher.Get(hashedKey)
		var tag entity.Tag
		err := json.Unmarshal([]byte(val.(string)), &tag)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "unexpected server error")
		}
		return c.JSON(http.StatusOK, dto.TagDtoRes{Message: "tag fetched successfully", Tag: &tag})
	}
	// retrieve from database
	tag, errApp := th.tagApp.GetTagById(id)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())

	}
	// put result in redis
	v, _ := json.Marshal(tag)
	err = th.cacher.Put(hashedKey, v, 600)
	if err != nil {
		c.Logger().Error("unable to put tag in redis cache")
	}
	return c.JSON(http.StatusOK, dto.TagDtoRes{Message: "tag fetched successfully", Tag: tag})
}

func (th *TagHandler) UpdateTag(c echo.Context) error {
	t := new(entity.Tag)
	if err := c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse request body")
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {

		return echo.NewHTTPError(http.StatusNotFound, "tag not found")
	}
	// search tag inside database
	tag, errApp := th.tagApp.GetTagById(id)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())

	}
	// modify its value
	tag.Tag = t.Tag
	// save again
	if _, errApp := th.tagApp.SaveTag(tag); err != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	// flush redis cache
	err = th.cacher.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}
	return c.JSON(http.StatusOK, dto.TagDtoRes{Message: "tag updated successfully", Tag: tag})
}

func (th *TagHandler) DeleteTag(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "tag not found")
	}
	tag, errApp := th.tagApp.GetTagById(id)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	errApp = th.tagApp.DeleteTag(tag.ID)
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	// flush redis cache
	err = th.cacher.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}
	return c.JSON(http.StatusOK, dto.TagDtoRes{Message: "tag deleted successfully"})
}
func (th *TagHandler) GetAllTag(c echo.Context) error {
	//check in redis
	hashedKey := utils.MD5([]byte("tags"))
	if th.cacher.IsExist(hashedKey) {
		fmt.Println("from redist")
		val := th.cacher.Get(hashedKey)
		var tags []entity.Tag
		err := json.Unmarshal([]byte(val.(string)), &tags)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "unexpected server error")
		}
		return c.JSON(http.StatusOK, dto.TagsDtoRes{Message: "tag fetched successfully", Tags: &tags})
	}
	// retrieve from database
	tags, errApp := th.tagApp.GetAllTag()
	if errApp != nil {
		return echo.NewHTTPError(errApp.Code, errApp.AsMessage())
	}
	// put the result in redis
	v, _ := json.Marshal(tags)
	err := th.cacher.Put(hashedKey, v, 600)
	if err != nil {
		c.Logger().Error("unable to put tag in redis cache")
	}
	return c.JSON(http.StatusOK, dto.TagsDtoRes{Message: "tags fetched successfully", Tags: &tags})
}
