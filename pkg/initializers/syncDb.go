package initializers

import "go-jwt/pkg/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
