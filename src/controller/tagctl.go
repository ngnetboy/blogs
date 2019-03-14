package controller

import (
	"net/http"
	"service"
	"utils"

	"github.com/gin-gonic/gin"
)

func GetTagAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	tag := service.Tag.GetTag()
	result.Data = tag
	return
}

func GetTagArticleAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	tagIdStr := c.Query("tag_id")
	if tagIdStr == "" {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		return
	}

	var tagId uint
	if err := utils.StrToUint(tagIdStr, &tagId); err != nil {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		return
	}

	result.Data = service.Tag.GetTagArticleByID(tagId)
	return
}
