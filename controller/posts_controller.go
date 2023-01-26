package controller

import (
	"github.com/MichiKaneko/nekoblog/db"
	"github.com/MichiKaneko/nekoblog/model"
	"github.com/gin-gonic/gin"
)

func GetPostById(c *gin.Context) {
	var post model.Post
	err := db.Database.First(&post, c.Param("id")).Error

	if err != nil {
		c.HTML(404, "404.html", gin.H{})
		return
	}

	if post == (model.Post{}) {
		c.HTML(404, "404.html", gin.H{})
		return
	}

	c.HTML(200, "show.html", gin.H{
		"post": post,
	})
}
