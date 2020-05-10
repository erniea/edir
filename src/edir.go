package main

import (
	"edir"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
)

func main() {
	os.Setenv("FYNE_FONT", "D2Coding.ttf")
	app := app.New()

	edirApp := edir.NewEdirApp()

	w := app.NewWindow("edir")
	w.Resize(fyne.NewSize(1024, 768))

	fileMenu := fyne.NewMenu("File", fyne.NewMenuItem("Quit", func() {}))
	d2Coding := func() {
		os.Setenv("FYNE_FONT", "D2Coding.ttf")
		app.Settings().SetTheme(theme.DarkTheme())
	}
	neoDgm := func() {
		os.Setenv("FYNE_FONT", "neodgm_code.ttf")
		app.Settings().SetTheme(theme.DarkTheme())
	}
	fontMenu := fyne.NewMenu("Font", fyne.NewMenuItem("D2Coding", d2Coding), fyne.NewMenuItem("NeoDGM", neoDgm))
	w.SetMainMenu(fyne.NewMainMenu(fileMenu, fontMenu))
	w.SetContent(edirApp.GetWidget())

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyLeft:
			edirApp.Tab.MoveLeft()
		case fyne.KeyDown:
			edirApp.Tab.MoveDown()
		case fyne.KeyUp:
			edirApp.Tab.MoveUp()
		case fyne.KeyRight:
			edirApp.Tab.MoveRight()
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
