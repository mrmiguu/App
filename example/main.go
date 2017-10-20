package main

import (
	"github.com/mrmiguu/app"
)

func main() {
	for i, url := range []string{
		"images/zcougar_dragonsun.png",
		"images/trsipic1_lazur.jpg",
		"images/alex-bisleys_horsy_5.png",
		"images/acryl_bobablast.png",
	} {
		go func(ld <-chan *app.Image, i int) {
			img := <-ld
			if img == nil {
				println("bad image url")
				return
			}

			// img.Resize(600-i*90, 600-i*90)
			// img.Show(true, 2500*time.Millisecond)

			// width, height := img.Size()

			// go img.Resize(-width, -height, 5000-i*250)
			// img.Move(width/2, height/2, 5000-i*250)
			// img.Show(false, 2500-i*125)
		}(app.LoadImage(url), i)
	}

	app.Serve()
}
