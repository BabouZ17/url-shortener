package model

import (
	"time"
)

type UrlCreationDTO struct {
	Target string `json:"target" binding:"required"`
}

type Url struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Alias     string    `json:"alias" gorm:"not null;unique"`
	Target    string    `json:"target" gorm:"not null;unique"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewUrl(alias, target string) Url {
	return Url{
		Alias:  alias,
		Target: target,
	}
}
