package component

import (
	"os"
	"strings"

	"fyne.io/fyne/widget"
)

//Element of folder
type Element struct {
	Name        string        // path
	Ext         string        // extension
	IsDirectory bool          // dir
	HBox        *widget.Box   // Hbox
	Indicator   *widget.Label // Selected Indicator
}

//NewElement New Element
func NewElement(fileInfo os.FileInfo) (elm *Element) {
	fullName := fileInfo.Name()
	if !fileInfo.IsDir() {
		dot := strings.LastIndex(fullName, ".")

		if dot >= 0 {
			elm = &Element{Name: fullName[:dot], Ext: fullName[dot+1:]}
		} else {
			elm = &Element{Name: fullName}
		}
	} else {
		elm = &Element{Name: fullName, IsDirectory: true}
	}

	return
}

//ToHBox to Hbox
func (elm *Element) ToHBox() *widget.Box {

	elm.Indicator = widget.NewLabel(" ")
	elm.HBox = widget.NewHBox(widget.NewLabel(elm.Name), elm.Indicator, widget.NewLabel(elm.Ext))
	return elm.HBox
}
