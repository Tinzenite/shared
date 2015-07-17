package shared

import (
	"fmt"
	"log"
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
func (v Version) Valid(that version, selfid string) (Version, bool) {
	if v.Max() > that.Max() {
		// local version is ahead
		log.Println("Local version is ahead of remote version!")
		return v, false
	}
	// if local changes don't even exist no need to check the following
	_, ok := v[selfid]
	if ok && v[selfid] != that[selfid] {
		// this means local version was changed without the other peer realizing
		log.Println("Merge conflict! Local file has since changed.")
		return v, false
	}
	// otherwise we can update
	return that, true
}

/*
String representation of version.
*/
func (v Version) String() string {
	var output string
	for key, value := range v {
		output += fmt.Sprintf("%s: %d\n", key, value)
	}
	return output
}
