package model

type User struct {
	Username  string `gorm:"type:varchar(256);not null;"`
	Password  string `gorm:"type:varchar(256);not null;"`
	Tag       string `gorm:"type:varchar(256);not null;"`
	Count     int `gorm:"type:int(255);not null;"`
}

func SelectUserByUserName(username string)User{
	var user User
	DB.Model(&User{}).Where("username = ?",username).First(&user)
	return user
}

func InsertUser(user User)error{
	return DB.Model(&User{}).Create(&user).Error
}

func UpdateTagByName(user User,tag string) error{
	return DB.Model(User{}).Where("username = ?",user.Username).Update("tag", tag).Error
}