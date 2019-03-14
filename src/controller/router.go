package controller

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	router := gin.Default()

	admin := router.Group("/api/v1/admin")
	{
		admin.GET("/category", GetCategoryAction)
		admin.POST("/category", AddCategoryAction)
		admin.DELETE("/category", DeleteCategoryAction)
		admin.PUT("/category", UpdateCategoryAction)
		admin.GET("/category/count", GetCategoryCountAction)
		admin.GET("/category/article", GetCategoryArticleAction)

		admin.GET("/article", GetArticleAction)
		admin.POST("/article", AddArticleAction)
		admin.PUT("/article", UpdateArticleAction)
		admin.DELETE("/article", DeleteArticleAction)
		admin.GET("/article/count", GetArticleCountAction)

		admin.GET("/tag", GetTagAction)
		admin.GET("/tag/article", GetTagArticleAction)
	}

	return router
}
