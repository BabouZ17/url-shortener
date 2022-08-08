package repository

import (
	"github.com/BabouZ17/url-shortener/pkg/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(id int) (model.Url, error)
	GetByAlias(alias string) (model.Url, error)
	List() ([]model.Url, error)
	Add(url model.Url) (model.Url, error)
	Delete(id string) error
	Count() (int64, error)
}

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{db}
}

func (repo *UrlRepository) Add(url model.Url) (model.Url, error) {
	if err := repo.db.Create(&url).Error; err != nil {
		return model.Url{}, err
	}
	return url, nil
}

func (repo *UrlRepository) List() ([]model.Url, error) {
	var urls []model.Url
	if err := repo.db.Find(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

func (repo *UrlRepository) GetById(id int) (model.Url, error) {
	var url model.Url

	if err := repo.db.First(&url, id).Error; err != nil {
		return model.Url{}, err
	}
	return url, nil
}

func (repo *UrlRepository) GetByAlias(alias string) (model.Url, error) {
	var url model.Url

	if err := repo.db.Where("alias = ?", alias).First(&url).Error; err != nil {
		return model.Url{}, err
	}
	return url, nil
}

func (repo *UrlRepository) Delete(id int) error {
	var url model.Url

	err := repo.db.Delete(&url, id)
	if err.Error != nil {
		return err.Error
	}

	if err.RowsAffected != 1 {
		return NewRepositoryError("record to delete not found")
	}

	return nil
}

func (repo *UrlRepository) Count() (int64, error) {
	var count int64
	var url model.Url

	if err := repo.db.Model(&url).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
