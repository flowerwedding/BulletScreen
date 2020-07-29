package service

import (
	"BulletScreen/model"
)

func FindUser(username string) bool{
	user := model.SelectUserByUserName(username)
	if user.Password != ""{
		return false
	}
	return true
}

func AddUser(user model.User) error{
	if err:= model.InsertUser(user);err != nil{
		return err
	}
	return nil
}

func CompareUserPassword(username string,password string) bool{
	user := model.SelectUserByUserName(username)
	if user.Password != password {
		return false
	}
	return true
}

func FindTag(username string) string{
	user := model.SelectUserByUserName(username)
	return user.Tag
}

func ChangeTag(name string,tag string)error{
	user := model.SelectUserByUserName(name)
	err :=model.UpdateTagByName(user,tag)
	return err
}