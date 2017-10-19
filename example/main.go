package main

import "github.com/mrmiguu/app"

func main() {
	imgld := app.LoadImage("images/mighty_no_09_cover_art_by_robduenas.jpgx")

	img := <-imgld
	if img == nil {
		println("bad image url")
	}
}
