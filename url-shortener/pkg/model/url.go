package model

type Url struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Alias  string `json:"alias" gorm:"not null"`
	Target string `json:"target" gorm:"not null"`
}
