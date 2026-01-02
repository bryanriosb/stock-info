package domain

import "time"

type RatingOption struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Label     string    `json:"label" gorm:"type:varchar(255);not null;uniqueIndex"`
	Value     string    `json:"value" gorm:"type:varchar(255);not null;uniqueIndex"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (RatingOption) TableName() string {
	return "rating_options"
}
