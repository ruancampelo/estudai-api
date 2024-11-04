package repository

import (
	. "estudai-api/internal/model"
	"gorm.io/gorm"
)

type ContentRepositoryGorm struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) *ContentRepositoryGorm {
	return &ContentRepositoryGorm{db: db}
}

func (r *ContentRepositoryGorm) Create(content *Content) error {
	return r.db.Create(content).Error
}

func (r *ContentRepositoryGorm) FindByID(id uint) (*Content, error) {
	var content Content
	if err := r.db.First(&content, id).Error; err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *ContentRepositoryGorm) FindAll() ([]Content, error) {
	var contents []Content
	if err := r.db.Find(&contents).Error; err != nil {
		return nil, err
	}
	return contents, nil
}

func (r *ContentRepositoryGorm) Update(content *Content) error {
	return r.db.Save(content).Error
}

func (r *ContentRepositoryGorm) Delete(id uint) error {
	return r.db.Delete(&Content{}, id).Error
}
