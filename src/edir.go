package main

import (
	"edir"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

func main() {
	app := app.New()

	edirApp := edir.NewEdirApp()

	w := app.NewWindow("Hello")
	w.SetContent(edirApp.GetWidget())
	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyUp:
			edirApp.Tab.MoveUp()
		case fyne.KeyDown:
			edirApp.Tab.MoveDown()
		case fyne.KeyEnter:
			fallthrough
		case fyne.KeyReturn:
			edirApp.Tab.MoveIn()
		case fyne.KeyBackspace:
			edirApp.Tab.MoveOut()
		}

	})
	w.ShowAndRun()
}
