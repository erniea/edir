package component

//PathTracker path tracker
type PathTracker struct {
	History []string //Path history

}

//Push History
func (pt *PathTracker) Push(path string) {
	pt.History = append(pt.History, path)
}

//Peek History
func (pt *PathTracker) Peek() (s string, last int) {
	last = len(pt.History) - 1
	if last >= 0 {
		s = pt.History[last]
	} else {
		s = ""
	}
	return
}

//Pop History
func (pt *PathTracker) Pop() (s string) {
	s, last := pt.Peek()
	pt.History = pt.History[:last-1]
	return
}
