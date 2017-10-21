package main

import (
	"sync"
	"time"

	"github.com/mrmiguu/app"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	for _, url := range []string{
		"images/zcougar_dragonsun.png",
		"images/trsipic1_lazur.jpg",
		"images/alex-bisleys_horsy_5.png",
		"images/acryl_bobablast.png",
	} {
		img, err := app.AddImage(url)
		if err != nil {
			panic(err)
		}

		// i := i
		go func() {
			defer wg.Done()

			time.Sleep(1000 * time.Millisecond)
			img.Show(true, 1*time.Second)
			// img.Resize(600-i*90, 600-i*90)

			// width, height := img.Size()

			// go img.Resize(-width, -height, 5000-i*250)
			// img.Move(width/2, height/2, 5000-i*250)
			// img.Show(false, 2500-i*125)
		}()
	}
	wg.Wait()

	app.Serve()
}
