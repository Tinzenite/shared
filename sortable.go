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

/*
SortableString is a string slice that can be sorted by length.
*/
type SortableString []*string

func (s SortableString) Len() int {
	return len(s)
}

func (s SortableString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableString) Less(i, j int) bool {
	// path are sorted alphabetically all by themselves! :D
	return *s[i] < *s[j]
}
