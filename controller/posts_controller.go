package controller

import (
	"github.com/MichiKaneko/nekoblog/db"
	"github.com/MichiKaneko/nekoblog/model"
	"github.com/gin-contrib/sessions"
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

func GetNewPost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")

	if username == nil {
		c.Redirect(302, "/user/login")
		return
	}

	user := getUserByUsername(username.(string))

	if user == (model.User{}) {
		c.Redirect(302, "/user/login")
		return
	}

	c.HTML(200, "new_post.html", gin.H{})
}

func CreateNewPost(c *gin.Context) {
	var post model.Post

	session := sessions.Default(c)
	username := session.Get("username")

	if username == nil {
		c.Redirect(302, "/user/login")
		return
	}

	user := getUserByUsername(username.(string))

	post.Title = c.PostForm("title")
	post.Content = c.PostForm("content")

	if post.Title == "" || post.Content == "" {
		c.HTML(404, "new_post.html", gin.H{
			"error": "Please fill in all fields",
		})
		return
	}

	post.UserID = user.ID

	if err := createPost(post.Title, post.Content, post.UserID); err != nil {
		c.HTML(404, "new_post.html", gin.H{
			"error": "Error creating post",
		})
		return
	}

	c.Redirect(302, "/user/posts")
}

func ShowUserPosts(c *gin.Context) {
	var posts []model.Post
	var user model.User

	session := sessions.Default(c)
	username := session.Get("username")

	if username == nil {
		c.Redirect(302, "/user/login")
		return
	}

	user = getUserByUsername(username.(string))

	if user == (model.User{}) {
		c.Redirect(302, "/user/login")
		return
	}

	db.Database.Where("user_id = ?", user.ID).Find(&posts)

	c.HTML(200, "user_posts.html", gin.H{
		"posts": posts,
	})
}

func createPost(title string, content string, userId uint) error {
	var post model.Post
	post.Title = title
	post.Content = content
	post.UserID = userId
	return db.Database.Create(&post).Error
}
