package service

import (
	"model"
	"sync"
)

var Category = &categoryService{
	mutex: &sync.Mutex{},
}

type categoryService struct {
	mutex *sync.Mutex
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
