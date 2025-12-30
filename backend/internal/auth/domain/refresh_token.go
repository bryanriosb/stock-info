package domain

import (
	"time"
)

type RefreshToken struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int64     `json:"user_id" gorm:"not null;index"`
	Token     string    `json:"-" gorm:"size:255;uniqueIndex;not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	Revoked   bool      `json:"revoked" gorm:"default:false"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

func (r *RefreshToken) IsValid() bool {
	return !r.Revoked && time.Now().Before(r.ExpiresAt)
}
