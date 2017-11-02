package util

import "fmt"

//https://github.com/fatih/color/blob/master/color.go
type Style byte
type Fg byte
type Bg byte

const 默认 = "0m"
const start = "\033["
const end = "\033[" + 默认
const (
	正常 Style = 0
	闪烁 = 5
	下划线 = 4
	粗线 = 1
)
const (
	前景黑 Fg = iota + 30
	前景红
	前景绿
	前景黄
	前景蓝
	前景紫红
	前景青蓝
	前景灰
	前景白
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