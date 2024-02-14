package user

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	resultSave := r.db.Create(&user)

	if resultSave.Error != nil {
		log.Fatal(resultSave.Error)
		return user, resultSave.Error
	}

	return user, nil
}

func (r *repository) FindEmail(email string) (User, error) {
	var user User
	resultFind := r.db.Where("email = ?", email).Find(&user)

	if resultFind.Error != nil {
		log.Fatal(resultFind.Error)
		return user, resultFind.Error
	}

	return user, nil
}
