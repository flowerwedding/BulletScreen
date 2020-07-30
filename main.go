package main

import (
	"BulletScreen/controller"
	"BulletScreen/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	//"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	r := gin.Default()
	m := melody.New()
	service.InitService()

	r.Use(cors.Default())

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session",store))

	r.POST("/BulletScreen/login", controller.Login)
	r.POST("/BulletScreen/register",controller.Register)

	for i := 1;i <= 3;i++ {
		go service.OpenConsumer()
	}


	r.GET("/BulletScreen/ws/:name",func(c *gin.Context) {
		controller.SendMessage(c,m)
	})

	r.GET("/BulletScreen/set/word",controller.SetWord)//设置敏感词
	r.POST("/BulletScreen/count/verify",controller.VerifyCount)//账户验证
	r.POST("/BulletScreen/blacklist/add",controller.AddBlacklist)//设置黑名单
	r.POST("/BulletScreen/price/take",controller.TakePrice)//红包抽奖
	r.POST("/BulletScreen/red/package",controller.MakeRed)//红包抽奖,红包和订单差不多，抢完就没

	controller.Do(m)

	_ = r.Run(":5000")
}