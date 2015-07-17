package shared

/*
Sortable allows sorting Objectinfos by path.
*/
type Sortable []*ObjectInfo

func (s Sortable) Len() int {
	return len(s)
}

func (s Sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sortable) Less(i, j int) bool {
	// path are sorted alphabetically all by themselves! :D
	return s[i].Path < s[j].Path
}
