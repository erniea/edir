package component

import (
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/widget"
)

//Tab tab
type Tab struct {
	pathTracker PathTracker   //Path Tracker
	vBox        *widget.Box   //TabBox
	PathLabel   *widget.Label //Label

	Elements          []*Element //Elements
	FocusedElementIdx int        //Focused
}

//NewTab to create
func NewTab(path string) (t *Tab) {
	t = &Tab{PathLabel: widget.NewLabel(path), vBox: widget.NewVBox()}
	t.Navigate(path)
	return
}

//GetWidget from tab
func (t *Tab) GetWidget() *widget.Box {
	return widget.NewVBox(t.PathLabel, t.vBox)
}

//Navigate to new path
func (t *Tab) Navigate(path string) {
	t.pathTracker.Push(path)
	t.PathLabel.SetText(path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	t.vBox.Children = t.vBox.Children[:0]
	t.Elements = t.Elements[:0]
	t.FocusedElementIdx = 0

	for _, file := range files {
		elm := NewElement(file)
		t.vBox.Append(elm.ToHBox())
		t.Elements = append(t.Elements, elm)
	}
	if len(t.Elements) > 0 {
		t.Elements[0].Indicator.SetText("-")
	}
	t.vBox.Refresh()

}

//MoveUp focused element
func (t *Tab) MoveUp() {

	if t.FocusedElementIdx > 0 {
		t.Elements[t.FocusedElementIdx].Indicator.SetText(" ")
		t.Elements[t.FocusedElementIdx-1].Indicator.SetText("-")
		t.FocusedElementIdx--
	}
}

//MoveDown focused element
func (t *Tab) MoveDown() {
	if t.FocusedElementIdx < len(t.Elements)-1 {
		t.Elements[t.FocusedElementIdx].Indicator.SetText(" ")
		t.Elements[t.FocusedElementIdx+1].Indicator.SetText("-")
		t.FocusedElementIdx++
	}
}

//MoveIn to focused element
func (t *Tab) MoveIn() {
	focusedElement := t.Elements[t.FocusedElementIdx]

	if focusedElement.IsDirectory {
		currentpath, _ := t.pathTracker.Peek()
		t.Navigate(currentpath + "\\" + focusedElement.Name)

	} else {

	}
}

//MoveOut from current path
func (t *Tab) MoveOut() {
	currentpath, _ := t.pathTracker.Peek()
	idx := strings.LastIndex(currentpath, "\\")
	if idx > 0 {
		t.Navigate(currentpath[:idx])
	}
}
