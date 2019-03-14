package service

import (
	"model"
	"sync"

	log "github.com/Sirupsen/logrus"
)

var Tag = &tagService{
	mutex: &sync.Mutex{},
}

type tagService struct {
	mutex *sync.Mutex
}

func (t *tagService) AddTag(tag *model.Tag) error {
	return db.Create(tag).Error
}
func (t *tagService) GetTagByName(name string) *model.Tag {
	var tag model.Tag

	if err := db.Where("name = ?", name).First(&tag).Error; err != nil {
		return nil
	}
	return &tag
}

func (t *tagService) GetTag() []*model.TagCount {
	var count []*model.TagCount

	if err := db.Select("name, count(name)").Group("name").Scan(&count).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Errorln("get category count error:", err.Error())
		}
		return nil
	}
	return count
}

func (t *tagService) GetTagArticleByID(tagId uint) []*model.ArticleBase {
	var articles []*model.ArticleBase

	if err := db.Table("articles").Where("id in(?)", db.Table("article_tag").Select("article_id").Where("tag_id=?", tagId)).Select("id, create_at, name, summary, amount").Scan(&articles).Error; err != nil {
		if gorm.ErrRecordNotFound != err {
			log.Errorln("get tag article by id error:", err.Error())
		}
		return nil
	}

	return articles
}
