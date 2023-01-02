package main

import (
	"embed"
	"fmt"
	"gouser32/win32"
	"gouser32/win32/w32"
	"image"
	_ "image/png"
	"math/rand"
	"time"
)

var (
	//go:embed assets/*
	assets embed.FS
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {
	opt := win32.OptionsWindow()
	opt.Dimension = image.Point{X: 600, Y: 400}
	opt.Position = image.Point{X: 10, Y: 10}
	opt.TitleName = "window"
	opt.Icon = GetImage("assets/icon.png")
	opt.Styles = w32.OVERLAPPED | w32.SYSMENU&^w32.SIZEBOX
	// opt.FullScreemStyle()
	if err := win32.Run(opt); err != nil {
		panic(err)
	}
	fmt.Println(opt)
}

func GetImage(filePath string) image.Image {
	file, err := assets.Open(filePath)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	return img
}
