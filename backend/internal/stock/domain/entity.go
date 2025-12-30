package domain

import "time"

type Stock struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Ticker     string    `json:"ticker" gorm:"size:10;not null;index"`
	Company    string    `json:"company" gorm:"size:255;not null;index"`
	Brokerage  string    `json:"brokerage" gorm:"size:255"`
	Action     string    `json:"action" gorm:"size:100"`
	RatingFrom string    `json:"rating_from" gorm:"size:50"`
	RatingTo   string    `json:"rating_to" gorm:"size:50"`
	TargetFrom float64   `json:"target_from" gorm:"type:decimal(10,2)"`
	TargetTo   float64   `json:"target_to" gorm:"type:decimal(10,2)"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}

func (Stock) TableName() string {
	return "stocks"
}
