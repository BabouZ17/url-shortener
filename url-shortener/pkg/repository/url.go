package repository

import (
	"gorm.io/gorm"
	"github.com/BabouZ17/url-shortener/pkg/model"
)

type Repository interface {
	Get(id string) (url model.Url, e error)
	List() (urls []model.Url)
	//Add(target string) (url model.Url)
	//Delete(id string)
	//DeleteAll()
}

type UrlRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UrlRepository {
	return &UrlRepository{db}
}

func (repo *UrlRepository) List() (urls []model.Url) {
	repo.db.Find(&urls)
	return urls
}

func (repo *UrlRepository) Get(id int) (url model.Url, e error) {
	result := repo.db.First(&url, id)
	if result.Error != nil {
		return model.Url{}, result.Error
	}
	return url, nil
}
