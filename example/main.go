package main

import "github.com/mrmiguu/app"

func main() {
	ld := app.LoadImage("images/mighty_no_09_cover_art_by_robduenas.jpg")

	img := <-ld
	if img == nil {
		println("bad image url")
	}

	select {}
}
