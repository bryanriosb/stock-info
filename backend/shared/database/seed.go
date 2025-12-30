package database

import (
	"log"

	"github.com/bryanriosb/stock-info/internal/user/domain"
	"github.com/bryanriosb/stock-info/shared"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB, cfg shared.AdminConfig) error {
	var existingAdmin domain.User
	result := db.Where("role = ?", domain.RoleAdmin).First(&existingAdmin)

	if result.Error == nil {
		log.Printf("Admin user already exists: %s", existingAdmin.Username)
		return nil
	}

	if result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	admin := &domain.User{
		Username: cfg.Username,
		Email:    cfg.Email,
		Role:     domain.RoleAdmin,
	}

	if err := admin.SetPassword(cfg.Password); err != nil {
		return err
	}

	if err := db.Create(admin).Error; err != nil {
		return err
	}

	log.Printf("Admin user created: %s", admin.Username)
	return nil
}
