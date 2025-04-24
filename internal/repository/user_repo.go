package repository

import (
	"errors"
	"fmt"
	rand "math/rand"
	"github.com/rmrachmanfauzan/bank_assessment/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user *model.User) error
	Find(id uint) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) RegisterUser(user *model.User) error {
	var count int64
	r.db.Model(&model.User{}).Where("nik = ?", user.NIK).Count(&count)
	if count > 0 {
		return errors.New("NIK already exists")
	}

	r.db.Model(&model.User{}).Where("phone = ?", user.Phone).Count(&count)
	if count > 0 {
		return errors.New("phone already exists")
	}
	
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	noRekening := fmt.Sprintf("%016d", rand.Intn(9999999999999999))

	account := model.Account{
		NoRekening: noRekening,
		UserID:     user.ID,
	}

	if err := r.db.Create(&account).Error; err != nil {
		return err
	}
	
	
	if err := r.db.Preload("Accounts").First(user, user.ID).Error; err != nil {
		return err
	}
	
	return nil
}

func (r *userRepository) Find(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.Preload("Accounts").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


