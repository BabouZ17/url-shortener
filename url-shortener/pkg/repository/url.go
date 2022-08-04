package repository

import "github.com/BabouZ17/url-shortener/pkg/model"

type UrlRepository interface {
	Get(id string)
	Add(target string) (url model.Url)
	List() (urls []model.Url)
	Delete(id string)
	DeleteAll()
}