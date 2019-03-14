package service

import (
	"model"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

var Category = &categoryService{
	mutex: &sync.Mutex{},
}

type categoryService struct {
	mutex *sync.Mutex
}

func (c *categoryService) GetCategoryCount() []*model.CategoryCount {
	var count []*model.CategoryCount

	if err := db.Select("name, count(name)").Group("name").Scan(&count).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Errorln("get category count error:", err.Error())
		}
		return nil
	}
	return count
}

func (c *categoryService) GetCategoryArticleByID(id uint) []*model.ArticleBase {
	var articles []*model.ArticleBase

	if err := db.Table("articles").Where("id in(?)", db.Table("article_category").Select("article_id").Where("category_id=?", id)).Select("id, create_at, name, summary, amount").Scan(&articles).Error; err != nil {
		if gorm.ErrRecordNotFound != err {
			log.Errorln("get tag article by id error:", err.Error())
		}
		return nil
	}

	return articles
}

func (c *categoryService) GetCategory() []*model.Category {
	var category []*model.Category

	if err := db.Find(&category).Error; err != nil {
		return nil
	}
	return category
}

func (c *categoryService) GetCategoryByName(name string) *model.Category {
	var category model.Category

	if err := db.Where("name = ?", name).First(&category).Error; err != nil {
		return nil
	}
	return &category
}

func (c *categoryService) GetCategoryByID(id uint) *model.Category {
	var category model.Category

	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil
	}
	return &category
}

func (c *categoryService) UpdateCategory(category *model.Category) error {
	return db.Model(category).Update(category).Error
}

func (c *categoryService) AddCategory(category *model.Category) error {
	return db.Create(category).Error
}

func (c *categoryService) DeleteCategory(category *model.Category) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = tx.Unscoped().Delete(category).Error
	if err != nil {
		return
	}

	err = tx.Model(category).Association("Article").Clear().Error
	if err != nil {
		return
	}

	return
}
