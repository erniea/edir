package edir

import (
	"component"

	"fyne.io/fyne/widget"
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
func (app *Edir) GetWidget() *widget.Box {
	return app.Tab.GetWidget()
}
