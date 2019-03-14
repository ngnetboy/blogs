package controller

import (
	"model"
	"net/http"
	"service"
	"utils"

	"github.com/gin-gonic/gin"
)

func GetCategoryAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	category := service.Category.GetCategory()

	result.Data = category
	return
}

func GetCategoryCountAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	result.Data = service.Category.GetCategoryCount()
	return
}

func GetCategoryArticleAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	categoryIdStr := c.Query("category_id")
	if categoryIdStr == "" {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		return
	}

	var categoryId uint
	if err := utils.StrToUint(categoryIdStr, &categoryId); err != nil {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		return
	}

	result.Data = service.Category.GetCategoryArticleByID(categoryId)
	return
}

func AddCategoryAction(c *gin.Context) {
	var category model.Category
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	if err := c.BindJSON(&category); err != nil {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		return
	}

	if category := service.Category.GetCategoryByName(category.Name); category != nil {
		result.Code = utils.ErrInternal
		result.Msg = "category has existed."
		return
	}

	if err := service.Category.AddCategory(&category); err != nil {
		result.Code = utils.ErrInternal
		result.Msg = "create category failed."
		return
	}
	return
}

func DeleteCategoryAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	type request struct {
		ID uint `json:"id"`
	}

	var req request

	if err := c.BindJSON(&req); err != nil {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		return
	}
	category := service.Category.GetCategoryByID(req.ID)
	if category == nil {
		return
	}

	if err := service.Category.DeleteCategory(category); err != nil {
		result.Code = utils.ErrInternal
		result.Msg = "delete category failed."
		return
	}
	return
}

func UpdateCategoryAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	type request struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Desc string `json:"description"`
	}

	var req request
	if err := c.BindJSON(&req); err != nil {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		return
	}

	category := service.Category.GetCategoryByID(req.ID)
	if category == nil {
		result.Code = utils.ErrInternal
		result.Msg = "can't find this category."
		return
	}

	category.Name = req.Name
	category.Description = req.Desc

	if err := service.Category.UpdateCategory(category); err != nil {
		result.Code = utils.ErrInternal
		result.Msg = "update category failed."
		return
	}
	return
}
