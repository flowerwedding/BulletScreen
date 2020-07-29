package util

import (
	"gopkg.in/gomail.v2"
	"math/rand"
)

func Email(to string) int{
	m := gomail.NewMessage()

	m.SetHeader("From", "2804991212@qq.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "BulletScreen")

	ran := rand.Intn(100)
	m.SetBody("text/html", "验证码：" + string(ran))

	d := gomail.NewDialer("smtp.qq.com", 587, "2804991212@qq.com", "xygdhlezsvirdebb")


	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return ran
}