package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	Token          string
	Created_at     time.Time
	Updated_at     time.Time
	DeletedAt      gorm.DeletedAt `gorm:"-"`
}
