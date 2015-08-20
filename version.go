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
	// special handling for selfid peer: must always be exactly the same
	localValue, localExist := v[selfid]
	remotValue, remotExist := that[selfid]
	// if it doesn't exist locally it may not exist on the other side
	if !localExist && remotExist {
		log.Println("Version: unsymmetric self existance of <"+selfid+">!", v, that)
		return false
	}
	// if it exists we must only make sure that all other values are ok
	if localExist && remotExist && localValue != remotValue {
		log.Println("Version: wrong value for self <"+selfid+">!", v, that)
		return false
	}
	// basically we need to guarantee that the other version has all updates we are aware of
	for localPeer, localValue := range v {
		// ignore selfid since we handled that before
		if localPeer == selfid {
			continue
		}
		thatValue, exists := that[localPeer]
		// make sure all peers we know of is known by the other version
		if !exists {
			log.Println("Version: missing update from <"+localPeer+">!", v, that)
			return false
		}
		// and if it knows of a peer, the version must be at least equal (may be higher in case we missed something)
		if localValue > thatValue {
			log.Println("Version: missing updates from <"+localPeer+">!", v, that)
			return false
		}
	}
	// if we reach this all values are legal
	return true
}

/*
Merge another version with this one. NOTE that it will only merge if listed peers
exist in v XOR that. The return value signifies whether that is the case - if
false, no changes to the version are merged.
*/
func (v Version) Merge(that Version) bool {
	for peer, value := range that {
		_, exist := v[peer]
		if exist {
			// log.Println("Version: merge failed because peer is in both versions!")
			return false
		}
		v[peer] = value
	}
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
