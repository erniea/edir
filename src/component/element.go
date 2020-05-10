package component

import (
	"os"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

//Element of folder
type Element struct {
	ToParent    bool
	Name        string // path
	Ext         string // extension
	IsDirectory bool   // dir
	FileSize    int64  //
	EditDate    time.Time

	Indicator *widget.Label // Selected Indicator
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
		elm.FileSize = fileInfo.Size()
		elm.EditDate = fileInfo.ModTime()
	} else {
		elm = &Element{Name: fullName, IsDirectory: true}
	}

	return
}

//GetWidget to Hbox
func (elm *Element) GetWidget(pageCount int) (hBox fyne.CanvasObject) {
	AdditionalInfo := "[DIR]"
	if !elm.IsDirectory {
		size := elm.FileSize
		sizeIdx := 0
		sizePostfix := []string{"", "KiB", "MiB", "GiB", "TiB"}
		for {
			if size < 1024 {
				AdditionalInfo = strconv.FormatInt(size, 10) + " " + sizePostfix[sizeIdx]
				break
			} else {
				size = size / 1024
				sizeIdx++
			}
		}
	}

	nameLabel := widget.NewLabel(elm.Name)
	elm.Indicator = widget.NewLabel(" ")
	extLabel := widget.NewLabel(elm.Ext)
	additionalLabel := widget.NewLabel(AdditionalInfo)

	additionalLabel.Alignment = fyne.TextAlignTrailing

	layout.NewHBoxLayout()
	hBox = widget.NewHBox(nameLabel, layout.NewSpacer(), elm.Indicator, extLabel, additionalLabel)
	return
}
