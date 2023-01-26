package controller

import (
	"github.com/MichiKaneko/nekoblog/db"
	"github.com/MichiKaneko/nekoblog/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAdmin(c *gin.Context) {
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

	c.HTML(200, "admin.html", gin.H{
		"user": user,
	})
}

func GetRegister(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}

func PostRegister(c *gin.Context) {
	var user model.User
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if username == "" || email == "" || password == "" {
		c.HTML(404, "register.html", gin.H{
			"error": "Please fill in all fields",
		})
		return
	}

	user = getUserByUsername(username)

	if user != (model.User{}) {
		c.HTML(404, "register.html", gin.H{
			"error": "User already exists",
		})
		return
	}

	if err := createUser(username, email, password); err != nil {
		c.HTML(404, "register.html", gin.H{})
		return
	}

	session := sessions.Default(c)
	session.Set("username", username)
	session.Save()

	c.Redirect(302, "/user/admin")
}

func GetLogin(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func PostLogin(c *gin.Context) {
	var user model.User
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.HTML(404, "login.html", gin.H{
			"error": "Please fill in all fields",
		})
		return
	}

	user = getUserByUsername(username)

	if user == (model.User{}) {
		c.HTML(404, "login.html", gin.H{
			"error": "User not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		c.HTML(404, "login.html", gin.H{
			"error": "Wrong password",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("username", username)
	session.Save()

	c.Redirect(302, "/user/admin")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(302, "/user/login")
}

func createUser(username, email, password string) []error {
	var user model.User

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return []error{err}
	}

	user.Username = username
	user.Email = email
	user.Password = string(passwordEncrypted)

	err = db.Database.Create(&user).Error

	if err != nil {
		return []error{err}
	}

	return nil
}

func getUserByUsername(username string) model.User {
	var user model.User
	db.Database.Where("username = ?", username).First(&user)
	return user
}
