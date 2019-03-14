package service

import (
	"model"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

var Article = &articleService{
	mutex: &sync.Mutex{},
}

type articleService struct {
	mutex *sync.Mutex
}

func (a *articleService) GetArticleByID(articleID uint) *model.Article {
	var article model.Article
	var categories []*model.Category

	if err := db.Where("id = ?", articleID).First(&article).Error; err != nil {
		return nil
	}

	if err := db.Model(&article).Related(&categories, "Category").Error; err != nil {
		return nil
	}
	return &article
}

func (a *articleService) GetArticle(page int) []*model.ArticleBase {
	var articles []*model.ArticleBase
	//var articles []*model.Article
	// if err := db.Offset(model.Conf.PageNum * page).Find(&articles).Error; err != nil {
	// 	return nil
	// }
	//if err := db.Find(&articles).Error; err != nil {
	if err := db.Table("articles").Select("id, create_at, name, summary, amount").Offset(model.Conf.PageNum * page).Scan(&articles).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Errorln("get article base error:", err.Error())
		}
		return nil
	}

	return articles
}

func (a *articleService) GetArticleCount() int {
	var count int

	if err := db.Model(&model.Article{}).Count(&count).Error; err != nil {
		count = -1
		return count
	}
	return count
}

func (a *articleService) UpdateArticle(article *model.Article) error {
	return db.Model(article).Updates(article).Error
}

func (a *articleService) AddArticle(article *model.Article) error {
	return db.Create(article).Error
}

func (a *articleService) DeleteArticle(article *model.Article) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = tx.Unscoped().Delete(article).Error
	if err != nil {
		return
	}

	err = tx.Model(article).Association("Category").Clear().Error
	if err != nil {
		return
	}

	err = tx.Model(article).Association("Tags").Clear().Error
	if err != nil {
		return
	}

	return
}
