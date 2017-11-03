package util

import "fmt"

//https://github.com/fatih/color/blob/master/color.go
//http://blog.csdn.net/codingwangfeng/article/details/6957079
type Style byte
type Fg byte
type Bg byte

const 默认 = "0m"
const start = "\033["
const end = "\033[" + 默认
const (
	正常 Style = 0
	黯淡
	加深
	斜体
	下划线
	闪烁
)
const (
	前景黑 Fg = iota + 30
	前景红 //31
	前景绿 //32
	前景黄 //33
	前景蓝 //34
	前景紫红 //35
	前景青蓝 //36
	前景白 //37
)
const (
	背景黑 Bg = iota + 40
	背景红
	背景绿
	背景黄
	背景蓝
	背景紫红
	背景青蓝
	背景白
)

type Color struct {
	style Style
	fg    Fg
	bg    Bg
}

func ColorNew(s Style, f Fg, b Bg) *Color {
	c := new(Color)
	c.style = s
	c.fg = f
	c.bg = b
	return c
}
func (c *Color)Out(s string) string {
	return fmt.Sprintf("%s%d;%d;%dm%s%s", start, c.style, c.fg, c.bg, s, end)
}