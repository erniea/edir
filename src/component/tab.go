package component

import (
	"io/ioutil"
	"log"
	"math"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

var Page = 20

//Tab tab
type Tab struct {
	pathTracker PathTracker //Path Tracker

	hScroll   *widget.ScrollContainer //Scroll
	hBox      *widget.Box             //HBox in scroll
	vBox      []*widget.Box           //VBox for list
	PathLabel *widget.Label           //Label

	Elements          []*Element //Elements
	FocusedElementIdx int        //Focused
}

//NewTab to create
func NewTab(path string) (t *Tab) {
	hBox := widget.NewHBox()
	hScroll := widget.NewHScrollContainer(hBox)

	t = &Tab{PathLabel: widget.NewLabel(path), hScroll: hScroll, hBox: hBox}
	t.Navigate(path)
	return
}

//GetWidget from tab
func (t *Tab) GetWidget() (vBox fyne.CanvasObject) {
	vBox = widget.NewVBox(t.PathLabel, t.hScroll)
	return
}

//Navigate to new path
func (t *Tab) Navigate(path string) {
	t.pathTracker.Push(path)
	t.PathLabel.SetText(path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	pageCount := int(math.Ceil(float64(len(files)) / float64(Page)))

	t.hBox.Children = t.hBox.Children[:0]
	t.Elements = t.Elements[:0]
	t.FocusedElementIdx = 0

	vBox := widget.NewVBox()
	elm := &Element{ToParent: true, Name: ".."}
	t.Elements = append(t.Elements, elm)
	vBox.Append(elm.GetWidget(pageCount))
	Idx := 1
	for _, file := range files {
		elm := NewElement(file)
		t.Elements = append(t.Elements, elm)
		vBox.Append(elm.GetWidget(pageCount))

		if Idx%Page == Page-1 {
			t.hBox.Append(vBox)
			vBox = widget.NewVBox()
		}
		Idx++
	}
	if len(vBox.Children) > 0 {
		t.hBox.Append(vBox)
	}
	if len(t.Elements) > 0 {
		t.Elements[0].Indicator.SetText("-")
	}

	t.hScroll.Refresh()
}

//MoveDec move
func (t *Tab) MoveDec(leap int) {

	if t.FocusedElementIdx > leap-1 {
		t.Elements[t.FocusedElementIdx].Indicator.SetText(" ")
		t.Elements[t.FocusedElementIdx-leap].Indicator.SetText("-")
		t.FocusedElementIdx -= leap
	}
}

//MoveInc
func (t *Tab) MoveInc(leap int) {
	if t.FocusedElementIdx < len(t.Elements)-leap {
		t.Elements[t.FocusedElementIdx].Indicator.SetText(" ")
		t.Elements[t.FocusedElementIdx+leap].Indicator.SetText("-")
		t.FocusedElementIdx += leap
	}

}

//MoveLeft focused element
func (t *Tab) MoveLeft() {
	t.MoveDec(Page)
}

//MoveDown focused element
func (t *Tab) MoveDown() {
	t.MoveInc(1)
}

//MoveUp focused element
func (t *Tab) MoveUp() {
	t.MoveDec(1)
}

//MoveRight focused element
func (t *Tab) MoveRight() {
	t.MoveInc(Page)
}

//MoveIn to focused element
func (t *Tab) MoveIn() {
	focusedElement := t.Elements[t.FocusedElementIdx]

	if focusedElement.ToParent {
		t.MoveOut()
	} else if focusedElement.IsDirectory {
		currentpath, _ := t.pathTracker.Peek()
		t.Navigate(currentpath + "\\" + focusedElement.Name)

	} else {

	}
}

//MoveOut from current path
func (t *Tab) MoveOut() {
	currentPath, _ := t.pathTracker.Peek()
	idx := strings.LastIndex(currentPath, "\\")
	if idx > 2 {
		t.Navigate(currentPath[:idx])
	} else if idx > 0 {
		t.Navigate(currentPath[:idx+1])
	}
}
