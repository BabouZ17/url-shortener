package model

type Url struct {
	Id string `json:"id"`
	Alias string `json:"alias"`
	Target string `json:"target" binding:"required"`
}