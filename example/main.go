package main

import (
	"sync"
	"time"

	"github.com/mrmiguu/app"
)

func init() {
	app.Width = 600
	app.Height = 800
}

func main() {
	// g := app.NewGroup(alex, mark)
	// img := g.Live.AddImage("mark-n-alex.png")
	// txt := g.Self.AddText("open")
	// img.OnShow(func() {

	// })
	// txt.OnHit(func() {
	// 	img.Show(true)
	// })

	var wg sync.WaitGroup
	wg.Add(1)
	for i, url := range []string{
		"images/mighty_no_09_cover_art_by_robduenas.jpg",
		// "images/cougar_dragonsun.png",
		// "images/trsipic1_lazur.jpg",
		// "images/archmage_in_your_face.png",
		// "images/acryl_bladerunner.png",
		// "images/acryl_bobablast.png",
		// "images/alex-bisleys_horsy_5.png",
	} {
		img, err := app.AddImage(url)
		if err != nil {
			panic(err)
		}

		i := i
		func() {
			defer wg.Done()

			img.Resize(600-i*90, 600-i*90)
			img.Show(true, 2500*time.Millisecond)

			width, height := img.Size()
			println(`width:`, width, `|height:`, height)

			go img.Resize(-width, -height, time.Duration(5000-i*250)*time.Millisecond)
			img.Move(width/2, height/2, time.Duration(5000-i*250)*time.Millisecond)
			img.Show(false, time.Duration(2500-i*125)*time.Millisecond)
		}()
	}
	wg.Wait()

	app.Serve()
}
