package model

import (
	"gorm.io/gorm"
)

type PostTag struct {
	gorm.Model
	PostID uint `json:"post_id"`
	TagID  uint `json:"tag_id"`
}
