package model

import (
	"PhotonTrail-backend/pkg/common"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name        string      `gorm:"type:varchar(255);not null" json:"name"`
	Email       string      `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password    string      `gorm:"type:varchar(255);not null" json:"password"`
	Role        common.Role `gorm:"type:int;not null;default:1" json:"role"`
	Avatar      string      `gorm:"type:varchar(1024)" json:"avatar"`
	Description string      `gorm:"type:varchar(1024)" json:"description"`
	Background  string      `gorm:"type:varchar(1024)" json:"background"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) Create(db *gorm.DB) (uint, error) {
	res := db.Create(&u)
	if res.Error != nil {
		return 0, res.Error
	}
	return u.ID, nil
}

func (u User) Update(db *gorm.DB, values map[string]interface{}) error {
	return db.Model(&User{}).Where("id = ?", u.Model.ID).Updates(values).Error
}

func (u User) Delete(db *gorm.DB) error {
	return db.Delete(&User{}, u.Model.ID).Error
}

func (u User) Get(db *gorm.DB) (*User, error) {
	var user User
	err := db.Where("id = ?", u.Model.ID).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &user, err
	}
	return &user, nil
}

func (u User) GetByUsername(db *gorm.DB) (*User, error) {
	var user User
	err := db.Where("name = ?", u.Name).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &user, err
	}
	return &user, nil
}

func (u User) GetByEmail(db *gorm.DB) (*User, error) {
	var user User
	err := db.Where("email = ?", u.Email).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &user, err
	}
	return &user, nil
}
