package controller

import (
	"BulletScreen/service"
	"github.com/gin-contrib/sessions"

	//	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AddBlacklist(c *gin.Context){
	name := c.PostForm("name")
	if err := service.ChangeTag(name,"用户已禁止");err != nil{
		c.JSON(200,gin.H{"statu":10001,"error":err})
		return
	}
	c.JSON(200,gin.H{"statu":10000,"message":"黑名单设置成功"})
}

func VerifyCount(c *gin.Context){
	session :=sessions.Default(c)
	username := session.Get("username").(string)

/*	email := c.PostForm("email")
	ran := util.Email(email)

	ticket := time.NewTicker(time.Minute)
	for t := range ticket.C{
		fmt.Println(t)
		cs := c.PostForm("ran")
		if cs == string(ran){
			ticket.Stop()
			break
		}
	}*/

    verify1 := c.PostForm("verify1")
	verify2 := c.PostForm("verify2")
	if verify1 == verify2{
		if err := service.ChangeTag(username,"用户可发弹幕");err != nil{
			c.JSON(200,gin.H{"statu":10001,"error":err})
			return
		}
	}
	c.JSON(200,gin.H{"statu":10000,"message":"用户可发弹幕"})
}