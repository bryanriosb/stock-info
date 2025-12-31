package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"size:50;uniqueIndex:idx_users_username;not null"`
	Email     string    `json:"email" gorm:"size:255;uniqueIndex:idx_users_email;not null"`
	Password  string    `json:"-" gorm:"size:255;not null"`
	Role      Role      `json:"role" gorm:"size:20;not null;default:'user'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (User) TableName() string {
	return "users"
}

func (u *User) SetPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
