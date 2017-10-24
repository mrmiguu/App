package main

func main() {
	menu := my.New("menu")
	menu.AddImage("for-me-only.png")

	world := our.New("world")
	world.AddImage("for-us-all.jpg")

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
