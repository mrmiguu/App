package main

import "github.com/mrmiguu/app"

func main() {
	img := app.LoadImage("images/acryl_bladerunner.png")
	if <-img == nil {
		println("bad image url")
	}
}
