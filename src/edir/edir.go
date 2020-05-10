package edir

import (
	"component"

	"fyne.io/fyne"
)

//Edir Main App
type Edir struct {
	//Tabs []*component.Tab
	Tab *component.Tab // tab
}

//NewEdirApp Creates New Edir App
func NewEdirApp() (edir *Edir) {

	edir = &Edir{Tab: component.NewTab(`C:\Workspace\go\edir`)}
	return
}

//GetWidget from app
func (app *Edir) GetWidget() fyne.CanvasObject {
	return app.Tab.GetWidget()
}
