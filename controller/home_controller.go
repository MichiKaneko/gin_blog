package controller

import (
	"github.com/MichiKaneko/nekoblog/db"
	"github.com/MichiKaneko/nekoblog/model"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	var posts []model.Post
	var navItems []model.NavItem
	db.Database.Find(&posts)
	db.Database.Find(&navItems)

	c.HTML(200, "index.html", gin.H{
		"title":    "Home",
		"posts":    posts,
		"navItems": navItems,
		"active":   "/",
	})
}
