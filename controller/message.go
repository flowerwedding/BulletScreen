package controller

import (
	"BulletScreen/service"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"strings"
	"time"
)

//本来用原生的websocket的，想把rabbitmq放在websocket里面，但是ws和declare的数据类型一直报错，然后...
func SendMessage(c *gin.Context,m *melody.Melody){
	color :=c.Query("color")
	user := c.Query("user")
	if color ==  "" {color = "black"}

	tag := service.FindTag(user)
	if tag == "用户已禁止"{
		return
	}
	//session :=sessions.Default(c)
	//user := session.Get("username").(string)

	Begin(m,color,user)

	_ = m.HandleRequest(c.Writer,c.Request)
}

func Begin(m *melody.Melody,color string,user string){
	m.HandleConnect(func(s *melody.Session) {
		s.Set("color",color)
		s.Set("user",user)
	})
}

func Do(m *melody.Melody){
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		color,_ := s.Get("color")
		user ,_ := s.Get("user")
		_ = m.BroadcastFilter(msg, func(q *melody.Session) bool {
			//msg = []byte(util.Color(color.(string), string(msg)))

			_, words := service.FindWord()
			for _,word := range words{
				if strings.Contains(string(msg), word) {
					return false
				}
			}

			_ = service.Message{Uid :user.(string),Message: string(msg), Color : color.(string),Time: time.Now().Format("2006-01-02 15:04:05")}
		/*	err := service.Order(message)
			if err != nil{
				return false
			}*/
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})
}

func SetWord(c *gin.Context){
	word := c.Query("word")
	if word == ""{
		c.JSON(200,gin.H{"error":"word is nil"})
	}
	err := service.SetWord(word)
	if err != nil{
		c.JSON(200,gin.H{"error":err})
	}
	c.JSON(200,gin.H{"message":"the word "+word+"is forbidden"})
}