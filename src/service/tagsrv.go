package service

import (
	"model"
	"sync"
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

func (t *tagService) GetTag() []*model.Tag {
	var tags []*model.Tag
	if err := db.Find(&tags).Error; err != nil {
		return nil
	}
	return tags
}
