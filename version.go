package shared

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

/*
Version implements a vector clock.
*/
type Version map[string]int

/*
Increase the version for the given peer based on the already existing versions.
*/
func (v Version) Increase(selfid string) {
	/*TODO catch overflow on version increase!*/
	v[selfid] = v.Max() + 1
}

/*
Max version number from all listed peers.
*/
func (v Version) Max() int {
	var max int
	for _, value := range v {
		if value >= max {
			max = value
		}
	}
	return max
}

/*
Valid checks whether the version can be automerged or whether manual resolution
is required.
*/
func (v Version) Valid(that Version, selfid string) bool {
	if v.Max() > that.Max() {
		// local version is ahead
		log.Println("Local version is ahead of remote version!")
		return false
	}
	// if local changes don't even exist no need to check the following
	_, ok := v[selfid]
	if ok && v[selfid] != that[selfid] {
		// this means local version was changed without the other peer realizing
		log.Println("Merge conflict! Local file has since changed.")
		return false
	}
	// otherwise we can update
	return true
}

/*
Equal checks whether the version per id match perfectly between the two.
*/
func (v Version) Equal(that Version) bool {
	// nil check
	if that == nil {
		return false
	}
	// length must be same
	if len(v) != len(that) {
		return false
	}
	// all entries must match
	for id, value := range v {
		thatValue, exists := that[id]
		if !exists || thatValue != value {
			return false
		}
	}
	// if everything runs through successfully, true
	return true
}

/*
String representation of version.
*/
func (v Version) String() string {
	var output string
	var values []string
	for key, value := range v {
		values = append(values, fmt.Sprintf("[%s:%d]", key, value))
	}
	// sort values for easy reading
	sortable := SortableString(values)
	sort.Sort(sortable)
	values = []string(sortable)
	// add identifier
	output = "Version{" + strings.Join(values, ",") + "}"
	return output
}
