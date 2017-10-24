package main

import (
	"github.com/mrmiguu/App/my"
	"github.com/mrmiguu/App/our"
)

func main() {
	menu := my.New("menu")
	menu.AddImage("myMenu.png")

	world := our.New("world")
	world.AddImage("ourMenu.jpg")

	start.Show()

	start.Keyboard = true

	start.HitFunc(func() {
		start.Hide()
		login.Show()
	})

	login.HitFunc(func() {
		login.Hide()
	})
}
