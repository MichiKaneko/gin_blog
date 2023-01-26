package controller

import (
	"github.com/MichiKaneko/nekoblog/db"
	"github.com/MichiKaneko/nekoblog/model"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	var posts []model.Post
	db.Database.Find(&posts)

	c.HTML(200, "index.html", gin.H{
		"posts": posts,
	})
}
