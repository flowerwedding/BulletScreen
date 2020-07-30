package controller

import (
	inits "BulletScreen/model"
	"BulletScreen/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login (c *gin.Context){//注册成功
	var user inits.User
	user.Username = c.PostForm("username")

	if !service.FindUser(user.Username){
		c.JSON(http.StatusOK, gin.H{"code": 10002, "message": "用户已注册"})
		return
	}

	user = inits.User{
		Password : c.PostForm("password"),
		Username : user.Username,
		Tag      : "账号未验证",
		Count    : 0,
	}

	if err := service.AddUser(user); err != nil{
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": "注册成功!"})
}

func Register (c *gin.Context) {//登录成功
	var user inits.User
	user.Username = c.PostForm("username")
	var password = c.PostForm("password")

	if !service.CompareUserPassword(user.Username,password){
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "密码错误"})
		return
	}

	session:=sessions.Default(c)
	session.Set("username", user.Username)
    _ = session.Save()

	tag := service.FindTag(user.Username)
	if tag != "用户已禁止"{
		tag = "账号未验证"
	}
	_ = service.ChangeTag(user.Username, tag)

	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": "登录成功！"})
}