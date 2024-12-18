package model

import (
	"gorm.io/gorm"
)

type Post struct {
	*gorm.Model
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	Camera  string `gorm:"type:varchar(255);" json:"camera"`
	Lens    string `gorm:"type:varchar(255)" json:"lens"`
	UserID  uint   `gorm:"type:int;not null" json:"user_id"`
}

type PostImage struct {
	*gorm.Model
	PostID uint   `gorm:"type:int;not null" json:"post_id"`
	UserID uint   `gorm:"type:int;not null" json:"user_id"`
	Url    string `gorm:"type:varchar(1024);not null" json:"url"`
	Width  int    `gorm:"type:int;not null" json:"width"`
	Height int    `gorm:"type:int;not null" json:"height"`
	Index  int    `gorm:"type:int;not null" json:"index"`
}

func (p Post) Get(db *gorm.DB, page, pageSize int) (*[]Post, error) {
	var posts []Post

	offset := (page - 1) * pageSize
	if err := db.Order("updated_at desc").Offset(offset).Limit(pageSize).Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func (p Post) GetByUserID(db *gorm.DB) (*[]Post, error) {
	var posts []Post

	if err := db.Where("user_id = ?", p.UserID).Order("updated_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}

	return &posts, nil
}

func (p Post) GetByID(db *gorm.DB) (*Post, error) {
	var post Post

	if err := db.First(&post, p.ID).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (p Post) GetImages(db *gorm.DB, postID uint) (*[]PostImage, error) {
	var images []PostImage

	if err := db.Where("post_id = ?", postID).Find(&images).Error; err != nil {
		return nil, err
	}

	return &images, nil
}

func (p Post) Create(db *gorm.DB, images *[]PostImage) (uint, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	for _, i := range *images {
		i.PostID = p.ID
		i.UserID = p.UserID
		if err := tx.Create(&i).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return p.ID, nil
}

func (p Post) Delete(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Delete(&Post{}, p.Model.ID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&PostImage{}, "post_id = ?", p.Model.ID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
