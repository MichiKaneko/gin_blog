package model

import (
	"gorm.io/gorm"
)

type NavItem struct {
	gorm.Model
	Name     string `json:"name"`
	Url      string `json:"url"`
	IsActive bool   `json:"is_active"`
}
