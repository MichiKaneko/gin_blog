package main

import (
	"os"

	"github.com/MichiKaneko/nekoblog/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	secret := os.Getenv("SECRET")
	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("nekoblog_user", store))

	r.LoadHTMLGlob("view/**/*")

	r.GET("/", controller.Home)
	r.GET("/posts/:id", controller.GetPostById)

	r.GET("/user/login", controller.GetLogin)
	r.POST("/user/login", controller.PostLogin)
	r.GET("/user/logout", controller.Logout)

	r.GET("/user/register", controller.GetRegister)
	r.POST("/user/register", controller.PostRegister)

	r.GET("/user/admin", controller.GetAdmin)

	r.GET("/user/posts", controller.ShowUserPosts)

	r.GET("/user/posts/new", controller.GetNewPost)
	r.POST("/user/posts/new", controller.CreateNewPost)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	return r
}
