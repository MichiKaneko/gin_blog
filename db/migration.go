package db

import (
	"log"

	"github.com/MichiKaneko/nekoblog/model"
)

func Migrate() {
	Database.AutoMigrate(&model.Post{})
	Database.AutoMigrate(&model.User{})
	Database.AutoMigrate(&model.NavItem{})
	log.Println("Database migrated")
}
