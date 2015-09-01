package shared

import (
	"fmt"
	"log"
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
Includes returns true if the version contains all knowledge and version of the
given version. Method is intended to allow ordering of versions.
*/
func (v Version) Includes(that Version) bool {
	for thatPeer, thatValue := range that {
		// must exist locally
		thisValue, thisExists := v[thatPeer]
		if !thisExists {
			return false
		}
		// and local must be equal or higher
		if thisValue < thatValue {
			return false
		}
	}
	return true
}

/*
Valid checks whether the version can be automerged or whether manual resolution
is required.

NOTE: The selfid value is a special case and ignored, so long as its knowledge
is correct. The value is however updated if the other version has a higher value.
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
		// See note.
		if thisPeer == selfid {
			// check if we need to update this version
			if thatValue > thisValue {
				// update self value
				log.Println("Version: WARNING: accepting update to self from other version!", v, that)
				v[selfid] = thatValue
			}
			// in any case do not further check the selfpeer
			continue
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
