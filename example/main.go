package main

func main() {
	login := app.AddText("Login")
	start := app.AddText("Start")

	start.Show()

	start.Keyboard = true

	start.HitFunc(func() {
		start.Hide()

		login.Show()
		login.HitFunc(func() {
			login.Hide()
		})
	})
}
