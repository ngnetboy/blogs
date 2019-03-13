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
