package service

import (
	"BulletScreen/model"
	"math/rand"
	"time"
)

type Message struct{
	Message string
	Color string
	Time string
	Uid  string
}

var i = 0

func AddMessage(message Message)error{
	 i ++
     bullet := model.Bullet{
		 Uid:     message.Uid,
		 Id:      i,
		 Context: message.Message,
		 Color:   message.Color,
		 Come:    " ",
		 Start:   message.Time,
		 Finish:  " ",
	 }

	 if err := model.AddBullet(bullet);err != nil{
	 	return err
	 }
	 return nil
}

func SetWord(word string) error{
	newword := model.Word{
		Word : word,
	}
	if err := model.InsertWord(newword) ;err != nil{
		return err
	}
	return nil
}

func FindWord() (error,[]string){
	var str []string
	words,err := model.SelectWord()
	if err != nil{
		return err,nil
	}
	for _,word := range words{
		str = append(str,word.Word)
	}
	return nil,str
}

func SelectMessage(ran int)(string,string,error){
	count ,err :=model.SelectNumCount()
	if err != nil{
		return "","",err
	}
	for ;;{
		if count > ran{
			break
		}
		rand.Seed(time.Now().UnixNano())
		ran = rand.Intn(5)
	}
	bullet := model.SelectBulletByCount(count - ran)
	return bullet.Uid,bullet.Context,nil
}