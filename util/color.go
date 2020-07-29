package util

import "fmt"

var colour = make(map[string]int)

const (
    black  = 30
    red    = 31
    green  = 32
    yellow = 33
    blue   = 34
    purple = 35
    cyan   = 36
    white  = 37
	underline = 4
)

func init(){
	colour["black"]  = black
	colour["red"]    = red
	colour["green"]  = green
	colour["yellow"] = yellow
	colour["blue"]   = blue
	colour["purple"]  = purple
	colour["cyan"]  = cyan
	colour["white"]  = white
}

//这个在控制台输出成功，但是在吗网页上一直乱码，加在websocket框架里面也这样
func Color(color string,s string) []byte {
	return  []byte(fmt.Sprintf("\033[0;;%dm%s\033[0m", colour[color],s))
}
