package model

type Bullet struct {
	Uid     string `gorm:"type:varchar(256);not null;"`
	Id      int `gorm:"type:int(255);not null;primary_key;"`
	Context string `gorm:"type:varchar(256);not null;"`
	Color   string `gorm:"type:varchar(256);not null;"`
	Come    string `gorm:"type:varchar(256);not null;"`
	Start   string `gorm:"type:varchar(256);not null;"`
	Finish  string `gorm:"type:varchar(256);not null;"`
}

type Word struct {
	Word  string `gorm:"type:varchar(256);not null;"`
}

func AddBullet(bullet Bullet) error{
	return DB.Model(&Bullet{}).Create(&bullet).Error
}

func InsertWord(word Word) error{
	return DB.Model(&Word{}).Create(&word).Error
}

func SelectWord()([]Word,error){
	var words []Word
	err := DB.Table("words").Find(&words).Error
	if err != nil {
		return nil, err
	}
	return words, nil
}

func SelectNumCount()(int, error){
	var bullets []Bullet
	var count int
	err :=DB.Table("bullets").Find(&bullets).Count(&count).Error
	if err != nil {return 0,err}
	return count,nil
}

func SelectBulletByCount(count int)Bullet{
	var bullet Bullet
	DB.Model(&Bullet{}).Where("id = ?",count).First(&bullet)
	return bullet
}