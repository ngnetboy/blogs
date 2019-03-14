package controller

import (
	"model"
	"net/http"
	"service"
	"strconv"
	"strings"
	"utils"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type articleRequest struct {
	ID       uint   `json:"article_id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Summary  string `json:"summary"`
	Draft    int    `json:"draft"`
	UserID   uint   `json:"user_id"`
	Category []uint `json:"category"`
	Tag      string `json:"tag"`
}

func GetArticleCountAction(c *gin.Context) {
	type articleResponse struct {
		Count   int `json:"count"`
		PageNum int `json:"page_num"`
	}

	var response articleResponse
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	response.Count = service.Article.GetArticleCount()
	response.PageNum = model.Conf.PageNum

	result.Data = response
	return
}

// GET /article?id=x&page=x
func GetArticleAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	idStr := c.Query("article_id")
	pageStr := c.Query("page")

	if idStr != "" {
		var articleID uint
		if err := utils.StrToUint(idStr, &articleID); err != nil {
			result.Code = utils.ErrInvalidArgument
			result.Msg = utils.ErrCodeMsg[result.Code]
			return
		}
		article := service.Article.GetArticleByID(articleID)
		if article != nil {
			article.Amount = article.Amount + 1
			service.Article.UpdateArticle(article)
		}
		result.Data = article
		return
	}

	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			result.Code = utils.ErrInvalidArgument
			result.Msg = utils.ErrCodeMsg[result.Code]
			return
		}
		article := service.Article.GetArticle(page)
		result.Data = article
		return
	}
}

func getCategoryByID(categoryIDs []uint) []*model.Category {
	var categories []*model.Category
	for _, categoryID := range categoryIDs {
		category := service.Category.GetCategoryByID(categoryID)
		if category != nil {
			categories = append(categories, category)
		}
	}
	return categories
}

func getTagByName(tagNames []string) []*model.Tag {
	var tags []*model.Tag
	for _, tagName := range tagNames {
		tag := service.Tag.GetTagByName(tagName)
		if tag == nil {
			tag = &model.Tag{Name: tagName}
			if err := service.Tag.AddTag(tag); err != nil {
				log.Errorln("add tag failed: ", err.Error())
				continue
			}
		}
		tags = append(tags, tag)
	}
	return tags
}

func AddArticleAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	var articlereq articleRequest
	if err := c.BindJSON(&articlereq); err != nil {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		log.Errorln("parse article argument error: ", err.Error())
		return
	}

	if len(articlereq.Summary) == 0 {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		return
	}
	// if user := service.User.GetUserByID(articlereq.UserID); user == nil {
	// 	result.Code = utils.ErrInternal
	// 	result.Msg = "invaild user."
	// 	return
	// }

	article := &model.Article{
		Name:    articlereq.Name,
		Content: articlereq.Content,
		Summary: articlereq.Summary,
		UserID:  articlereq.UserID,
		Draft:   articlereq.Draft,
	}

	article.Category = getCategoryByID(articlereq.Category)
	article.Tags = getTagByName(strings.Split(articlereq.Tag, ","))

	if err := service.Article.AddArticle(article); err != nil {
		result.Code = utils.ErrInternal
		result.Msg = "add article failed"
		return
	}
	return
}

func UpdateArticleAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	var articlereq articleRequest
	if err := c.BindJSON(&articlereq); err != nil {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		return
	}

	article := service.Article.GetArticleByID(articlereq.ID)
	if article == nil {
		result.Code = utils.ErrInternal
		result.Msg = "can't find this article"
		return
	}

	article.Name = articlereq.Name
	article.Content = articlereq.Content
	article.Summary = articlereq.Summary
	article.Draft = articlereq.Draft

	article.Category = getCategoryByID(articlereq.Category)
	article.Tags = getTagByName(strings.Split(articlereq.Tag, ","))

	if err := service.Article.UpdateArticle(article); err != nil {
		result.Code = utils.ErrInternal
		result.Msg = "add article failed"
		return
	}
	return
}

func DeleteArticleAction(c *gin.Context) {
	result := utils.NewResult()
	defer c.JSON(http.StatusOK, result)

	var articlereq articleRequest
	if err := c.BindJSON(&articlereq); err != nil {
		result.Code = utils.ErrInvalidArgument
		result.Msg = utils.ErrCodeMsg[result.Code]
		return
	}

	article := service.Article.GetArticleByID(articlereq.ID)
	if article == nil {
		return
	}

	if err := service.Article.DeleteArticle(article); err != nil {
		result.Code = utils.ErrInternal
		result.Msg = "delete article failed"
		return
	}
	return
}
