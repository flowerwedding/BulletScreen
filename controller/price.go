package controller

import (
	"BulletScreen/service"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func TakePrice(c *gin.Context){
	rand.Seed(time.Now().UnixNano())
	ran := rand.Intn(5)

	uid,context,err := service.SelectMessage(ran)
	if err != nil{
		c.JSON(200,gin.H{"statu":10001,"error":err})
	}

	c.JSON(200,gin.H{"user":uid,"message":context})
}