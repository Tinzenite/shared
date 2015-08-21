package shared

import (
	"fmt"
	"sort"
	"strings"
)

/*
Version implements a vector clock. The value of zero should never turn up.
*/
type Version map[string]int

/*
CreateVersion returns a new Version object
*/
func CreateVersion() Version {
	return Version{}
}

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
	// for every value this knows of...
	for thisPeer, thisValue := range v {
		// ... that must have an entry (otherwise missing updates)
		thatValue, thatExists := that[thisPeer]
		if !thatExists {
			// log.Println("Version: failed knowledge check:", v, that)
			return false
		}
		// and that's value must be => than the local one
		if !(thatValue >= thisValue) {
			// log.Println("Version: failed value check:", v, that)
			return false
		}
		// that's it! if those two are right the version is valid!
	}
	// log.Println("Version: legal:", v, that)
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

/*
IsEmpty returns whether any entries have been made in this version.
*/
func (v Version) IsEmpty() bool {
	return len(v) == 0
}
